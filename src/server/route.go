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
	t, err := template.ParseFiles("template/html/404.html")
	if err != nil {
		logger.Error(err.Error())
	}

	t.Execute(w, nil)
}

func userLoginHandler(w http.ResponseWriter, r *http.Request) {
	s := manager.SessionStart(w, r)
	t, err := template.ParseFiles("template/html/user_login.html")
	if err != nil {
		logger.Error(err.Error())
	}

	pageData := newPageData(w, r, s)
	t.Execute(w, pageData)
}

func userLogoutHandler(w http.ResponseWriter, r *http.Request) {
	manager.SessionDestroy(w, r)

	t, err := template.ParseFiles("template/html/user_login.html")
	if err != nil {
		logger.Error(err.Error())
	}

	pageData := newPageData(w, r, nil)
	t.Execute(w, pageData)
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
	s := manager.SessionStart(w, r)
	if s.Get("is_login") != "1" {
		http.Redirect(w, r, "/user_login", 302)
		return
	}

	handler := uploadFileAPI{session:s}
	handler.handle(w, r)
}

func listFileAPIHandler(w http.ResponseWriter, r *http.Request) {
	s := manager.SessionStart(w, r)
	handler := listFileAPI{session:s}
	handler.handle(w, r)
}

func editFileHandler(w http.ResponseWriter, r *http.Request) {
	s := manager.SessionStart(w, r)

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
		http.Redirect(w, r, "/not_found", 302)
		return
	}


	pageData := newPageData(w, r, s)

	data.pageData = pageData
	t.Execute(w, data)
}

func deleteFileAPIHandler(w http.ResponseWriter, r *http.Request) {
	s := manager.SessionStart(w, r)
	if s.Get("is_login") != "1" {
		http.Redirect(w, r, "/user_login", 302)
		return
	}

	handler := deleteFileAPI{session:s}
	handler.handle(w, r)
}

func userLoginAPIHandler(w http.ResponseWriter, r *http.Request) {
	s := manager.SessionStart(w, r)
	handler := userLoginAPI{session:s}
	handler.handle(w, r)
}