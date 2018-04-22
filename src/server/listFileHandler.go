package main

import (
	"net/http"
	"html/template"
)

func listFileHandler(w http.ResponseWriter, r *http.Request)  {
	s := manager.SessionStart(w, r)

	t, err := template.ParseFiles("template/html/list_file.html")
	if err != nil {
		logger.Error(err.Error())
	}

	data := newPageData(w, r, s)
	t.Execute(w, data)
}
