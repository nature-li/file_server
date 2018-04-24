package main

import (
	"net/http"
	"html/template"
	"net/url"
	"path/filepath"
	"session"
	"strconv"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(filepath.Join(config.privateTemplatePath, "html/404.html"))
	if err != nil {
		logger.Error(err.Error())
	}

	t.Execute(w, nil)
}

func notAllowHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(filepath.Join(config.privateTemplatePath, "html/refuse.html"))
	if err != nil {
		logger.Error(err.Error())
	}

	t.Execute(w, nil)
}

func userLoginHandler(w http.ResponseWriter, r *http.Request) {
	s := manager.SessionStart(w, r)
	t, err := template.ParseFiles(filepath.Join(config.publicTemplatePath, "html/user_login.html"))
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
	// 检测是否登录
	if !checkLogin(s) {
		http.Redirect(w, r, "/user_login", 302)
		return
	}

	// 检测上传权限
	if !checkRight(s, UPLOAD_RIGHT) {
		http.Redirect(w, r, "/not_allowed", 302)
		return
	}

	handler := uploadFileAPI{session:s}
	handler.handle(w, r)
}

func listFileAPIHandler(w http.ResponseWriter, r *http.Request) {
	s := manager.SessionStart(w, r)
	if !checkLogin(s) {
		http.Redirect(w, r, "/user_login", 302)
		return
	}

	if !checkRight(s, DOWNLOAD_RIGHT) {
		http.Redirect(w, r, "/not_allowed", 302)
		return
	}

	handler := listFileAPI{session:s}
	handler.handle(w, r)
}

func editFileHandler(w http.ResponseWriter, r *http.Request) {
	s := manager.SessionStart(w, r)
	if !checkLogin(s) {
		http.Redirect(w, r, "/user_login", 302)
		return
	}

	if !checkRight(s, DOWNLOAD_RIGHT) {
		http.Redirect(w, r, "/not_allowed", 302)
		return
	}

	err := r.ParseForm()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	queryId := r.Form.Get("id")

	t, err := template.ParseFiles(filepath.Join(config.privateTemplatePath, "html/edit_file.html"))
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
	if !checkLogin(s) {
		http.Redirect(w, r, "/user_login", 302)
		return
	}

	if !checkRight(s, UPLOAD_RIGHT) {
		http.Redirect(w, r, "/not_allowed", 302)
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
	http.ServeFile(w, r, config.publicTemplatePath + "/img/favicon.ico")
}

func captchaAPIHandler(w http.ResponseWriter, r *http.Request) {
	s := manager.SessionStart(w, r)
	handler := newCaptchaHandler(s)
	handler.handle(w, r)
}

func downloadFileHandler(w http.ResponseWriter, r *http.Request) {
	// 检测登录
	s := manager.SessionStart(w, r)
	if !checkLogin(s) {
		http.Redirect(w, r, "/user_login", 302)
		return
	}

	// 检测上传权限
	if !checkRight(s, DOWNLOAD_RIGHT) {
		http.Redirect(w, r, "/not_allowed", 302)
		return
	}

	dataFs := http.FileServer(http.Dir(config.UploadDataPath))
	http.StripPrefix("/data/", dataFs).ServeHTTP(w, r)
}

func privateFileHandler(w http.ResponseWriter, r *http.Request) {
	// 检测登录
	s := manager.SessionStart(w, r)
	if !checkLogin(s) {
		http.Redirect(w, r, "/user_login", 302)
		return
	}

	// 检测上传权限
	if !checkRight(s, DOWNLOAD_RIGHT) {
		http.Redirect(w, r, "/not_allowed", 302)
		return
	}

	dataFs := http.FileServer(http.Dir(config.privateTemplatePath))
	http.StripPrefix("/templates/private/", dataFs).ServeHTTP(w, r)
}

func checkLogin(s session.Session) bool {
	isLogin := s.Get("is_login")
	if isLogin == "1" {
		return true
	}

	return false
}

func checkRight(s session.Session, right int64) bool {
	var userRight string
	if userRight = s.Get("user_right"); userRight == "" {
		return false
	}
	digitRight, err := strconv.ParseInt(userRight, 10, 64)
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	if digitRight & right == 0 {
		return false
	}

	return true
}