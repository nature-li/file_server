package main

import (
	"net/http"
	"html/template"
	"fmt"
)

func listFileHandler(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	wrapperClass := ""
	ret := r.Form.Get("toggle")
	if ret == "true" {
		wrapperClass = "toggled"
	}

	pinLock := "glyphicon-pushpin"
	ret = r.Form.Get("lock")
	if ret == "true" {
		pinLock = "glyphicon-lock"
	}

	t, err := template.ParseFiles("template/html/list_file.html")
	if err != nil {
		logger.Error(err.Error())
	}

	fmt.Println(wrapperClass)
	fmt.Println(pinLock)
	data := pageData{IsLogin:true, LoginName:"lyg", WrapperClass:wrapperClass, PinLock:pinLock}
	t.Execute(w, data)
}
