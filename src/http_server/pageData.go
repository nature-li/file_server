package main

import (
	"net/http"
	"session"
)

type pageData struct {
	IsLogin           bool
	LoginName         string
	WrapperClass      string
	PinLock           string
	HiddenClass       string
	UploadMaxFileSize int64
	UploadMaxFileSizeStr string
}

func newPageData(w http.ResponseWriter, r *http.Request, s session.Session) *pageData {
	// 是否展开侧边栏
	wrapperClass := ""
	hiddenClass := ""
	if cookie, ok := r.Cookie("pin_nav"); ok == nil {
		if cookie.Value == "1" {
			wrapperClass = "toggled"
			hiddenClass = "hidden-self"
		}
	}

	// 是否锁住浮动锁
	pinLock := "glyphicon-pushpin"
	if cookie, ok := r.Cookie("pin_lock"); ok == nil {
		if cookie.Value == "1" {
			pinLock = "glyphicon-lock"
		}
	}

	// 登录相关
	var isLogin = false
	var loginName = ""
	if s != nil {
		if s.Get("is_login") == "1" {
			isLogin = true
			loginName = s.Get("user_name")
		}
	}

	return &pageData{IsLogin: isLogin, LoginName: loginName, WrapperClass: wrapperClass, PinLock: pinLock, HiddenClass: hiddenClass, UploadMaxFileSize:config.UploadMaxSize, UploadMaxFileSizeStr: config.maxUploadSizeStr}
}