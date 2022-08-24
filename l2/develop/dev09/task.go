package main

import (
	"bytes"
	"flag"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type flags struct {
	A bool
}

func main() {
	var fl flags
	flag.BoolVar(&fl.A, "A", true, "Download all website")
	flag.Parse()
	//url := flag.Args()[0]
	url := "https://go.dev"
	wget(fl, url)
}

//получаем все сылки и асеты для страницы
func getLinks(domain string, depth int) (page Page, attachments []string, err error) {
	page.Links = make(map[string]int, 1000)
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

	var f func(*html.Node)
	f = func(n *html.Node) {
		// Get CSS
		if n.Type == html.ElementNode && n.Data == "link" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					link, err := resp.Request.URL.Parse(a.Val)
					if err == nil {
						attachments = append(attachments, link.String())
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
			page.Links[newLink] = depth
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return
}

type Page struct {
	depth int
	URL   string
	Links map[string]int
	HTML  string
}

func wget(fl flags, url string) {
	if fl.A {
		resp, err := http.Get(url)
		if err != nil {
			log.Println("Domain could not be reached!")
			return
		}
		defer resp.Body.Close()

		var files []string

		newLinks := make(chan map[string]int, 100000) // New links to scan
		pages := make(chan Page, 100000)              // Pages scanned
		attachments := make(chan []string, 100000)    // Attachments
		started := make(chan int, 100000)             // Crawls started
		finished := make(chan int, 100000)            // Crawls finished

		seen := make(map[string]bool)
		saved := make(map[string]Page)

		page, attached, _ := getLinks(url, 0)
		pages <- page
		attachments <- attached

		// Save links
		newLinks <- page.Links
		seen[url] = true

	EXIT:
		for {
			select {
			case links := <-newLinks:
				for link, depth := range links {
					if depth+1 > 2 {
						continue
					}
					if !strings.Contains(link, url) {
						continue
					}
					if !seen[link] || link[len(link)-1] == '/' && !seen[link[:len(link)-1]] { //проверить, что это наш домен, проверить на слэш в конце
						started <- 1
						seen[link] = true
						page, attached, _ = getLinks(link, depth+1)
						pages <- page
						attachments <- attached
						newLinks <- page.Links
						finished <- 1
					}
				}
			case page2 := <-pages:
				if _, ok := saved[page2.URL]; !ok {
					saved[page2.URL] = page2
					SaveHTML(page2.URL, page2.HTML, resp.Request.URL.String())
				}
			case attachment := <-attachments:
				for _, link := range attachment {
					SaveAttachment(link, resp.Request.URL.String())
					files = append(files, link)
				}
			default:
				if len(newLinks) == 0 && len(pages) == 0 && len(attachments) == 0 {
					log.Println("\nFinished scraping the site...")
					break EXIT //goto EXIT2
				}
			}
			//if len(started) > 1 && len(scanning) == 0 && len(started) == len(finished) {
			//	break
			//}
		}

		//log.Println("\nFinished scraping the site...")
	} else { //скачка файла, а не всего сайта
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
	//EXIT2:
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func SaveHTML(url string, html, root string) (err error) {
	if !exists("dir") {
		os.MkdirAll("dir", 0755) // first create directory
	}
	os.Chdir("dir")
	defer os.Chdir("C:\\Users\\los28\\GolandProjects\\L0\\l2\\develop\\dev09")

	filepath := strings.Replace(url, root, "", 1)
	if filepath == "" || filepath == "/" {
		filepath = "index.html"
	}
	lastSlash := strings.LastIndex(filepath, "/")
	if lastSlash != -1 {
		dir := filepath[1 : lastSlash+1]
		if !exists(dir) {
			os.MkdirAll(dir, 0755) // first create directory
		}
		os.Chdir(dir)
		filepath = filepath[lastSlash+1:]
	}
	//dir := filepath
	//if strings.HasSuffix(dir, ".html") {
	//	dir = dir[:len(dir)-5]
	//}
	//
	//if !exists(dir) {
	//	os.MkdirAll(dir, 0755) // first create directory
	//}

	//os.Chdir(dir)

	//str := url[8:]
	str := filepath
	if !strings.Contains(str, ".html") {
		str = str + ".html"
	}
	f, err := os.Create(str)
	if err != nil {

		return
	}
	defer f.Close()

	_, err = io.Copy(f, bytes.NewBufferString(html))

	return
}

func SaveAttachment(url, root string) (err error) {
	if !exists("dir") {
		os.MkdirAll("dir", 0755) // first create directory
	}
	os.Chdir("dir")
	defer os.Chdir("C:\\Users\\los28\\GolandProjects\\L0\\l2\\develop\\dev09")
	filepath := strings.Replace(url, root, "", 1)
	if filepath == "" {
		return
	}

	// Get last path

	lastSlash := strings.LastIndex(filepath, "/")
	if lastSlash != -1 {
		dir := filepath[1 : lastSlash+1]
		if !exists(dir) {
			os.MkdirAll(dir, 0755) // first create directory
		}
		os.Chdir(dir)
		filepath = filepath[lastSlash+1:]
	}

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	f, err := os.Create(filepath)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return
}
