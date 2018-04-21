package main

import (
	"net/http"
)

type pageData struct {
	IsLogin           bool
	LoginName         string
	WrapperClass      string
	PinLock           string
	HiddenClass       string
	UploadMaxFileSize string
}

func newPageData(w http.ResponseWriter, r *http.Request, isLogin bool, loginName string) *pageData {
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

	return &pageData{IsLogin: isLogin, LoginName: loginName, WrapperClass: wrapperClass, PinLock: pinLock, HiddenClass: hiddenClass, UploadMaxFileSize: maxUploadSizeStr}
}