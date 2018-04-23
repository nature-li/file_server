package main

import (
	"net/http"
	"html/template"
	"path/filepath"
)

func listFileHandler(w http.ResponseWriter, r *http.Request)  {
	s := manager.SessionStart(w, r)


	t, err := template.ParseFiles(filepath.Join(config.HttpTemplatePath, "html/list_file.html"))
	if err != nil {
		logger.Error(err.Error())
	}

	data := newPageData(w, r, s)
	t.Execute(w, data)
}
