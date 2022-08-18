package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type flags struct {
	A bool
}

func main() {
	var fl flags
	flag.BoolVar(&fl.A, "A", false, "Download all website")
	flag.Parse()
	url := flag.Args()[0]
	//var url string
	//fmt.Scan(&url)
	//req, _ := http.Get(url)
	//defer req.Body.Close()
	wget(fl, url)
}

func wget(fl flags, url string) {
	if fl.A {
		req, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer req.Body.Close()

		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatal(err)
		}
		err = os.Mkdir("downloadedWebsite", 0777)
		if err != nil {
			panic(err)
		}
		err = os.Chdir("downloadedWebsite")
		if err != nil {
			log.Fatal(err)
		}
		defer os.Chdir("..")
		file, err := os.Create("downloadWebsite" + time.Now().Format("15040502012006") + ".html")

		if err != nil {
			fmt.Println("Unable to create file:", err)
			os.Exit(1)
		}
		defer file.Close()
		file.WriteString(string(b))
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
