package main

import (
	"net/http"
	"fmt"
	"mtlog"
	"server/session/cookie"
)

func main() {
	logger = mtlog.NewLogger(false, mtlog.DEVELOP, mtlog.INFO, mtLogPath, mtLogName, mtLogMaxFileSize, mtLogKeepFileCount)
	if !logger.Start() {
		fmt.Println("logger.Start failed")
	}

	var err error
	manager, err = cookie.NewManager(cookieSecret, cookieSessionId, cookieAccessTime, sessionAliveTime)
	if err != nil {
		logger.Error("NewManager failed")
		return
	}

	fs := http.FileServer(http.Dir(httpTemplatePath))
	http.Handle("/template/", http.StripPrefix("/template/", fs))

	dataFs := http.FileServer(http.Dir(httpDataPath))
	http.Handle("/data/", http.StripPrefix("/data/", dataFs))

	http.HandleFunc("/", listFileHandler)
	// 上传文件
	http.HandleFunc("/upload_file", uploadFileHandler)
	http.HandleFunc("/upload_file_api", uploadFileAPIHandler)
	// 文件列表
	http.HandleFunc("/list_file", listFileHandler)
	http.HandleFunc("/file_list_api", listFileAPIHandler)
	// 编辑文件
	http.HandleFunc("/edit_file", editFileHandler)
	// 删除文件
	http.HandleFunc("/delete_file_api", deleteFileAPIHandler)
	// 404页面
	http.HandleFunc("/not_found", notFoundHandler)
	// 登录页面
	http.HandleFunc("/user_login", userLoginHandler)
	http.HandleFunc("/user_login_api", userLoginAPIHandler)
	// OA登录
	http.HandleFunc("/user_login_auth", userLoginAuthHandler)
	http.HandleFunc("/user_login_auth_api", userLoginAuthAPIHandler)
	// 退出登录
	http.HandleFunc("/user_logout", userLogoutHandler)

	err = http.ListenAndServe(serverListPort, nil)
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Stop()
}
