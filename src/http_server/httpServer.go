package main

import (
	"net/http"
	"fmt"
	"mtlog"
	"session/cookie"
	"flag"
	"session"
)

var logger *mtlog.Logger
var manager session.Manager
var config conf

func main() {
	var confPath = flag.String("conf", "", "config file path")
	flag.Parse()

	if *confPath == "" {
		fmt.Println("conf is empty")
		return
	}

	if config.getConf(*confPath) == nil {
		fmt.Println("parse config file error")
		return
	}
	config.show()

	logger = mtlog.NewLogger(false, mtlog.DEVELOP, mtlog.INFO, config.LogPath, config.LogName, config.LogFileSize, config.LogFileCount)
	if !logger.Start() {
		fmt.Println("logger.Start failed")
	}

	var err error
	manager, err = cookie.NewManager(config.HttpCookieSecret, config.HttpSessionId, config.HttpAccessTime, config.HttpSessionTimeout)
	if err != nil {
		logger.Error("NewManager failed")
		return
	}

	// 公开模板文件
	publicFs := http.FileServer(http.Dir(config.publicTemplatePath))
	http.Handle("/templates/public/", http.StripPrefix("/templates/public/", publicFs))
	// 私有模板文件
	http.HandleFunc("/templates/private/", privateFileHandler)
	// 图标
	http.HandleFunc("/favicon.ico", faviconHandler)

	// 404页面
	http.HandleFunc("/not_found", notFoundHandler)
	// 拒绝访问页面
	http.HandleFunc("/not_allowed", notAllowHandler)
	// 验证码
	http.HandleFunc("/captcha", captchaAPIHandler)

	// 登录页面
	http.HandleFunc("/user_login", userLoginHandler)
	http.HandleFunc("/user_login_api", userLoginAPIHandler)
	// OA登录
	http.HandleFunc("/user_login_auth", userLoginAuthHandler)
	http.HandleFunc("/user_login_auth_api", userLoginAuthAPIHandler)
	// 退出登录
	http.HandleFunc("/user_logout", userLogoutHandler)

	// 下载文件
	http.HandleFunc("/data/", downloadFileHandler)
	// 首页
	http.HandleFunc("/", listFileHandler)
	// 上传文件
	http.HandleFunc("/upload_file", uploadFileHandler)
	// 文件列表
	http.HandleFunc("/list_file", listFileHandler)
	// 编辑文件
	http.HandleFunc("/edit_file", editFileHandler)
	// 用户列表
	http.HandleFunc("/list_user", listUserHandler)

	// 上传文件
	http.HandleFunc("/upload_file_api", uploadFileAPIHandler)
	// 删除文件
	http.HandleFunc("/delete_file_api", deleteFileAPIHandler)
	// 文件列表
	http.HandleFunc("/file_list_api", listFileAPIHandler)
	// 用户列表
	http.HandleFunc("/list_user_api", listUserAPIHandler)
	// 编辑用户
	http.HandleFunc("/edit_user_api", editUserAPIHandler)


	err = http.ListenAndServe(config.HttpListenPort, nil)
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Stop()
}
