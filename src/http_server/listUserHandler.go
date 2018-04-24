package main

import (
	"net/http"
	"path/filepath"
	"html/template"
)

func listUserHandler(w http.ResponseWriter, r *http.Request)  {
	s := manager.SessionStart(w, r)
	if !checkLogin(s) {
		if config.ServerLocalMode {
			http.Redirect(w, r, "/user_login", 302)
		} else {
			http.Redirect(w, r, "/user_login_auth", 302)
		}
		return
	}

	if !checkRight(s, MANAGER_RIGHT) {
		http.Redirect(w, r, "/not_allowed", 302)
		return
	}

	t, err := template.ParseFiles(filepath.Join(config.privateTemplatePath, "html/list_user.html"))
	if err != nil {
		logger.Error(err.Error())
	}

	data := newPageData(w, r, s)
	t.Execute(w, data)
}
