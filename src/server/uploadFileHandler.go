package main

import (
	"net/http"
	"html/template"
	"strconv"
)

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	s := manager.SessionStart(w, r)
	if s.Get("is_login") != "1" {
		http.Redirect(w, r, "/user_login", 302)
		return
	}

	t, err := template.ParseFiles("template/html/upload_file.html")
	if err != nil {
		logger.Error(err.Error())
	}

	cookie := http.Cookie{Name: "upload_max_file_limit", Value: strconv.FormatInt(maxUploadSize, 10), Path: "/", HttpOnly: true, MaxAge: 0}
	http.SetCookie(w, &cookie)

	data := newPageData(w, r)
	t.Execute(w, data)
}