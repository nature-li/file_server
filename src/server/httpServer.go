package main

import (
	"net/http"
	"fmt"
	"mtlog"
	"server/session"
	"server/session/cookie"
)

// hello world, the web server
var logger *mtlog.Logger
var manager session.Manager

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
	logger = mtlog.NewLogger(false, mtlog.DEVELOP, mtlog.INFO, "./logs", "server", 100*1024*1024, -1)
	if !logger.Start() {
		fmt.Println("logger.Start failed")
	}

	fs := http.FileServer(http.Dir("./template"))
	http.Handle("/template/", http.StripPrefix("/template/", fs))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/learn", learnHandler)
	http.HandleFunc("/query_file_list", queryFileListHandler)

	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		logger.Fatalf("ListenAndServe: ", err)
	}

	logger.Stop()
}
