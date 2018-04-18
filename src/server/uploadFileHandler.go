package main

import (
	"net/http"
	"html/template"
)

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/html/upload_file.html")
	if err != nil {
		logger.Error(err.Error())
	}
	data := newPageData(w, r, true, "lyg")
	t.Execute(w, data)
}