package main

import (
	"net/http"
	"html/template"
)

func listFileHandler(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	wrapperClass := r.Form.Get("toggle")
	pinLock := r.Form.Get("lock")
	hiddenClass := r.Form.Get("hidden")

	t, err := template.ParseFiles("template/html/list_file.html")
	if err != nil {
		logger.Error(err.Error())
	}

	data := pageData{IsLogin:true, LoginName:"lyg", WrapperClass:wrapperClass, PinLock:pinLock, HiddenClass:hiddenClass}
	t.Execute(w, data)
}
