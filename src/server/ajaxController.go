package main

import (
	"net/http"
	_ "github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/thrsafe"
	"encoding/json"
)

type Result struct{
	Ret int
	Reason string
	Data interface{}
}

type ajaxController struct {
}

func (o *ajaxController)LoginAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	err := r.ParseForm()
	if err != nil {
		OutputJson(w, 0, "参数错误", nil)
		return
	}

	admin_name := r.FormValue("admin_name")
	admin_password := r.FormValue("admin_password")

	if admin_name == "" || admin_password == ""{
		OutputJson(w, 0, "参数错误", nil)
		return
	}

	// 存入cookie,使用cookie存储
	cookie := http.Cookie{Name: "admin_name", Value: admin_name, Path: "/"}
	http.SetCookie(w, &cookie)

	OutputJson(w, 1, "操作成功", nil)
	return
}

func OutputJson(w http.ResponseWriter, ret int, reason string, i interface{}) {
	out := &Result{ret, reason, i}
	b, err := json.Marshal(out)
	if err != nil {
		return
	}
	w.Write(b)
}