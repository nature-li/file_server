package main

import (
	"net/http"
	"html/template"
	"net/url"
	"path/filepath"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(filepath.Join(config.HttpTemplatePath, "html/404.html"))
	if err != nil {
		logger.Error(err.Error())
	}

	t.Execute(w, nil)
}

func userLoginHandler(w http.ResponseWriter, r *http.Request) {
	s := manager.SessionStart(w, r)
	t, err := template.ParseFiles(filepath.Join(config.HttpTemplatePath, "html/user_login.html"))
	if err != nil {
		logger.Error(err.Error())
		return
	}

	handler := newCaptchaHandler(s)
	_, hex, err := handler.createCaptcha()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	value := struct {
		Base64Img string
		*pageData
	} {
		Base64Img: hex,
		pageData : newPageData(w, r, s),
	}

	t.Execute(w, value)
}

func userLogoutHandler(w http.ResponseWriter, r *http.Request) {
	manager.SessionDestroy(w, r)

	http.Redirect(w, r, "/user_login", 302)
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

	t, err := template.ParseFiles(filepath.Join(config.HttpTemplatePath, "html/edit_file.html"))
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
	var redirectUrl = config.OauthAuthUrl
	redirectUrl += "?appid=" + config.OauthAppId
	redirectUrl += "&response_type=code"
	redirectUrl += "&redirect_uri=" + url.QueryEscape(config.OauthRedirectUrl)
	redirectUrl += "&scope=user_info"
	redirectUrl += "&state=test"
	http.Redirect(w, r, redirectUrl, 302)
}

func userLoginAuthAPIHandler(w http.ResponseWriter, r *http.Request) {
	s := manager.SessionStart(w, r)
	handler := userLoginAuthAPI{session:s}
	handler.handle(w, r)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, config.HttpTemplatePath + "/img/favicon.ico")
}

func captchaAPIHandler(w http.ResponseWriter, r *http.Request) {
	s := manager.SessionStart(w, r)
	handler := newCaptchaHandler(s)
	handler.handle(w, r)
}