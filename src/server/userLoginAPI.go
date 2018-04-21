package main

import (
	"net/http"
	"database/sql"
	"encoding/json"
	"crypto/md5"
	"encoding/hex"
	"strings"
	_ "server/session/cookie"
)

type userLoginAPI struct {
	Success bool `json:"success"`
	Msg string `json:"message"`

	db *sql.DB
}

func (o *userLoginAPI) handle(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	userName := r.Form.Get("user_name")
	password := r.Form.Get("user_password")

	if userName == "" {
		logger.Error("userName is empty")
		o.render(w, false, "USER_NAME_EMPTY")
		return
	}

	if password == "" {
		logger.Error("password is empty")
		o.render(w, false, "PASSWORD_EMPTY")
		return
	}

	o.db, err = sql.Open("sqlite3", sqliteDbPath)
	if err != nil {
		logger.Error(err.Error())
		o.render(w, false, "OPEN_DB_FAILED")
		return
	}
	defer o.db.Close()

	success, message := o.checkPassword(userName, password)
	if success == true {
		s := manager.SessionStart(w, r)
		s.Set("is_login", "1")
		s.Set("user_name", userName)
	}
	o.render(w, success, message)
}

func (o *userLoginAPI) render(w http.ResponseWriter, success bool, desc string) {
	o.Success = success
	o.Msg = desc

	result, err := json.Marshal(o)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	_, err = w.Write(result)
	if err != nil {
		logger.Error(err.Error())
	}
}

func (o *userLoginAPI) checkPassword(name, password string) (success bool, message string) {
	md5Value := md5.Sum([]byte(password))
	hexMd5 := hex.EncodeToString(md5Value[:])
	querySql := "SELECT user_name,passwd FROM user_list WHERE user_name=?"
	rows, err := o.db.Query(querySql, name)
	if err != nil {
		logger.Error(err.Error())
		return false, "QUERY_DB_FAILED"
	}

	var nameInDB string
	var passwdInDB string
	var count = 0
	for rows.Next() {
		err = rows.Scan(&nameInDB, &passwdInDB)
		if err != nil {
			logger.Error(err.Error())
			return false, "SCAN_DB_FAILED"
		}

		count++
	}

	if count == 0 {
		return false, "USER_NOT_EXISTS"
	}

	if !strings.EqualFold(hexMd5, passwdInDB) {
		return false, "PASSWORD_ERROR"
	}

	return true, "SUCCESS"
}
