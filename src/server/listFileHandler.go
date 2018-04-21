package main

import (
	"net/http"
	"html/template"
	"fmt"
)

func listFileHandler(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("template/html/list_file.html")
	if err != nil {
		logger.Error(err.Error())
	}

	s := manager.SessionStart(w, r)
	fmt.Println("user_name:", s.Get("user_name"))

	data := newPageData(w, r, true, "lyg")
	t.Execute(w, data)
}
