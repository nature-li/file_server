package main

import (
	"net/http"
	"html/template"
	"net/url"
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

func userLoginAuthHandler(w http.ResponseWriter, r *http.Request) {
	var redirectUrl = ""
	redirectUrl += "?appid=" + serverAuthAppId
	redirectUrl += "&response_type=code"
	redirectUrl += "&redirect_uri=" + url.QueryEscape(serverAuthRedirectUrl)
	redirectUrl += "&scope=user_info"
	redirectUrl += "&state=test"
	http.Redirect(w, r, redirectUrl, 302)
}