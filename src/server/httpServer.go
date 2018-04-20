package main

import (
	"net/http"
	"fmt"
	"mtlog"
	"server/session/cookie"
)

func init() {
	key := "1234567890"
	key += "1234567890"
	key += "1234567890"
	key += "12"

	var err error
	//manager, err = memory.NewManager("session_id", 60)
	manager, err = cookie.NewManager(key, "session_id", "last_access_time", 60)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	s := manager.SessionStart(w, r)

	name := s.Get("name")
	if name == nil {
		s.Set("name", "lyg")
	}
	fmt.Println("name:", name)
	fmt.Fprintln(w, "login")
}

func logout(w http.ResponseWriter, r *http.Request) {
	manager.SessionDestroy(w, r)
	fmt.Fprintln(w, "logout")
}

func delete(w http.ResponseWriter, r *http.Request) {
	s := manager.SessionStart(w, r)

	s.Del("name")
	fmt.Fprintln(w, "deleted")
}

func main() {
	logger = mtlog.NewLogger(false, mtlog.DEVELOP, mtlog.INFO, mtLogPath, mtLogName, mtLogMaxFileSize, mtLogKeepFileCount)
	if !logger.Start() {
		fmt.Println("logger.Start failed")
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

	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/learn", learnHandler)

	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		logger.Fatalf("ListenAndServe: ", err)
	}

	logger.Stop()
}
