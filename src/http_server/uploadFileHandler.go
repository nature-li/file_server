package main

import (
	"net/http"
	"html/template"
	"path/filepath"
)

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	s := manager.SessionStart(w, r)
	if s.Get("is_login") != "1" {
		http.Redirect(w, r, "/user_login", 302)
		return
	}

	t, err := template.ParseFiles(filepath.Join(config.HttpTemplatePath, "html/upload_file.html"))
	if err != nil {
		logger.Error(err.Error())
	}

	data := newPageData(w, r, s)
	t.Execute(w, data)
}