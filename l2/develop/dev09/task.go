package main

import (
	"bytes"
	"flag"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type flags struct {
	A bool
}

var (
	extensions = []string{".png", ".jpg", ".jpeg", ".json", ".js", ".pdf", ".txt", ".gif", ".bmp", ".zip", ".svg", ".json", ".xml", ".dll", ".css"}
	validURL   = regexp.MustCompile(`\(([^()]*)\)`)
	validCSS   = regexp.MustCompile(`\{(\s*?.*?)*?\}`)
)

func main() {
	var fl flags
	flag.BoolVar(&fl.A, "A", true, "Download all website")
	flag.Parse()
	//url := flag.Args()[0]
	url := "https://artemonweb2.herokuapp.com"
	wget(fl, url)
}

func getLinks(domain string) (page Page, attachments []string, err error) {
	resp, err := http.Get(domain)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	page.HTML = buf.String()

	doc, err := html.Parse(buf)
	if err != nil {
		log.Println(err)
		return
	}

	page.URL = domain

	//foundMeta := false

	var f func(*html.Node)
	f = func(n *html.Node) {
		for _, a := range n.Attr {
			if a.Key == "style" {
				if strings.Contains(a.Val, "url(") {
					found := string(validURL.Find([]byte(a.Val)))
					if found != "" {
						link, err := resp.Request.URL.Parse(found)
						if err == nil {
							attachments = append(attachments, link.String())
						}
					}
				}
			}
		}

		//if n.Type == html.ElementNode && n.Data == "meta" {
		//	for _, a := range n.Attr {
		//		if a.Key == "name" && a.Val == "robots" {
		//			foundMeta = true
		//		}
		//		if foundMeta {
		//			if a.Key == "content" && strings.Contains(a.Val, "noindex") {
		//				page.NoIndex = true
		//			}
		//		}
		//	}
		//}

		// Get CSS and AMP
		if n.Type == html.ElementNode && n.Data == "link" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					link, err := resp.Request.URL.Parse(a.Val)
					if err == nil {
						page.Links = append(page.Links, link.String())
					}
				}
			}
		}

		// Get JS Scripts
		if n.Type == html.ElementNode && n.Data == "script" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					link, err := resp.Request.URL.Parse(a.Val)
					if err == nil {
						attachments = append(attachments, link.String())
					}
				}
			}
		}

		// Get Images
		if n.Type == html.ElementNode && n.Data == "img" {
			for _, a := range n.Attr {
				if a.Key == "src" {
					link, err := resp.Request.URL.Parse(a.Val)
					if err == nil {
						attachments = append(attachments, link.String())
					}
				}
				if a.Key == "srcset" {
					links := strings.Split(a.Val, " ")
					for _, val := range links {
						link, err := resp.Request.URL.Parse(val)
						if err == nil {
							attachments = append(attachments, link.String())
						}
					}
				}
			}
		}

		// Get links
		if n.Type == html.ElementNode && n.Data == "a" {
			var newLink string
			for _, a := range n.Attr {
				if a.Key == "href" {
					link, err := resp.Request.URL.Parse(a.Val)
					if err == nil {
						newLink = link.String()
					}
				}

			}
			page.Links = append(page.Links, newLink)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return
}

type Page struct {
	URL       string
	Canonical string
	Links     []string
	NoIndex   bool
	HTML      string
}

func wget(fl flags, url string) {
	if fl.A {
		resp, err := http.Get(url)
		if err != nil {
			log.Println("Domain could not be reached!")
			return
		}
		defer resp.Body.Close()
		//Path - ./website
		//Domain - url

		var indexed, files []string

		scanning := make(chan int, 3)              // Semaphore
		newLinks := make(chan []string, 100000)    // New links to scan
		pages := make(chan Page, 100000)           // Pages scanned
		attachments := make(chan []string, 100000) // Attachments
		started := make(chan int, 100000)          // Crawls started
		finished := make(chan int, 100000)         // Crawls finished

		seen := make(map[string]bool)

		page, attached, _ := getLinks(url)
		pages <- page
		attachments <- attached

		// Save links
		newLinks <- page.Links
		seen[url] = true

		for {
			select {
			case links := <-newLinks:
				for _, link := range links {
					if !seen[link] {
						seen[link] = true
						page, attached, _ = getLinks(link)
						pages <- page
						attachments <- attached
					}
				}
			case page := <-pages:
				indexed = append(indexed, page.URL)
				go SaveHTML(page.URL, page.HTML, resp.Request.URL.String())

			case attachment := <-attachments:
				for _, link := range attachment {
					files = append(files, link)
				}
			}

			if len(started) > 1 && len(scanning) == 0 && len(started) == len(finished) {
				break
			}
		}

		log.Println("\nFinished scraping the site...")
	} else {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		downloadedFile := strings.Split(url, "/")
		out, err := os.Create(downloadedFile[len(downloadedFile)-1])
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		io.Copy(out, resp.Body)
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func SaveHTML(url string, html, root string) (err error) {
	filepath := strings.Replace(url, root, "", 1)
	if filepath == "" {
		filepath = "/index.html"
	}
	//
	//if len(strings.Split(url, "/")) > 1 {
	//	if string(filepath[len(filepath)-1]) == "/" {
	//		// if the url is a final url in a folder, like example.com/path
	//		// this will create the folder "path" and, inside, the index.html file
	//		if !exists(url + filepath) {
	//			os.MkdirAll(url+filepath, 0755) // first create directory
	//			filepath = filepath + "index.html"
	//		}
	//	} else {
	//		// if the url is not a final url in a folder, like example.com/path/bum.html
	//		// this will create the folder "path" and, inside, the bum.html file
	//		paths := strings.Split(url, "/")
	//		var path string
	//		if len(paths) <= 1 {
	//			path = url
	//		} else {
	//			total := paths[:len(paths)-1]
	//			path = strings.Join(total[:], "/")
	//
	//		}
	//		if !exists(url + path) {
	//			os.MkdirAll(url+path, 0755) // first create directory
	//		}
	//	}
	//}

	if !exists("dir") {
		os.MkdirAll("dir", 0755) // first create directory
	}

	os.Chdir("dir")
	defer os.Chdir("..")
	str := url[8:] + "index.html"
	f, err := os.Create(str)
	if err != nil {

		return
	}
	defer f.Close()

	_, err = io.Copy(f, bytes.NewBufferString(html))

	return
}
