package main

import (
	"net/http"
	"html/template"
	"net"
	"fmt"
	"io"
	"time"
	"strconv"
)

type IndexPageData struct {
	IsLogin bool
	LoginName string
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
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

func uploadFileAPIHandler(w http.ResponseWriter, r *http.Request) {
	handler := uploadFileAPI{}
	handler.handle(w, r)
}

func listFileAPIHandler(w http.ResponseWriter, r *http.Request) {
	handler := listFileAPI{}
	handler.handle(w, r)
}

func editFileHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	queryId := r.Form.Get("id")

	t, err := template.ParseFiles("template/html/edit_file.html")
	if err != nil {
		logger.Error(err.Error())
		return
	}

	handler := editFileAPI{}
	data := handler.queryFileList(queryId)
	if data == nil {
		return
	}

	if data.Id == 0 {
		http.Redirect(w, r, "/not_found", 301)
		return
	}

	cookie := http.Cookie{Name: "upload_max_file_limit", Value: strconv.FormatInt(maxUploadSize, 10), Path: "/", HttpOnly: true, MaxAge: 0}
	http.SetCookie(w, &cookie)

	pageData := newPageData(w, r, true, "lyg")
	pageData.setPageToggle(w)

	data.pageData = pageData
	t.Execute(w, data)
}

func deleteFileAPIHandler(w http.ResponseWriter, r *http.Request) {
	handler := deleteFileAPI{}
	handler.handle(w, r)
}