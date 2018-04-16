package main

import (
	"net/http"
	"html/template"
)

type loginController struct {
}

func (o *loginController)IndexAction(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/html/login/index.html")
	if err != nil {
		logger.Error(err.Error())
	}
	t.Execute(w, nil)
}
