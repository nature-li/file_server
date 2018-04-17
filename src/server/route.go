package main

import (
	"net/http"
	"html/template"
	"net"
	"fmt"
	"io"
	"time"
)

type IndexPageData struct {
	IsLogin bool
	LoginName string
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/login/index", http.StatusFound)
	}

	t, err := template.ParseFiles("template/html/404.html")
	if err != nil {
		logger.Error(err.Error())
	}

	t.Execute(w, nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/html/index.html")
	if err != nil {
		logger.Error(err.Error())
	}
	data := IndexPageData{IsLogin:true, LoginName:"lyg"}
	t.Execute(w, data)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/html/test.html")
	if err != nil {
		logger.Error(err.Error())
	}
	data := IndexPageData{IsLogin:true, LoginName:"lyg"}
	t.Execute(w, data)
}

func learnHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/html/learn.html")
	if err != nil {
		logger.Error(err.Error())
	}
	data := IndexPageData{IsLogin:true, LoginName:"lyg"}
	t.Execute(w, data)
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	url := "https://www.python.org/ftp/python/2.7.14/Python-2.7.14.tgz"

	timeout := time.Duration(5) * time.Second
	transport := &http.Transport{
		ResponseHeaderTimeout: timeout,
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, timeout)
		},
		DisableKeepAlives: true,
	}
	client := &http.Client{
		Transport: transport,
	}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	//copy the relevant headers. If you want to preserve the downloaded file name, extract it with go's url parser.
	w.Header().Set("Content-Disposition", "attachment; filename=Wiki.png")
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))

	//stream the body to the client without fully loading it into memory
	io.Copy(w, resp.Body)
}