package main

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
	"database/sql"
	"session"
)

type userLoginAuthAPI struct {
	AccessToken string `json:"access_token"`
	OpenId      string `json:"openid"`

	session session.Session
}

func (o *userLoginAuthAPI)handle(w http.ResponseWriter, r *http.Request) {
	s := manager.SessionStart(w, r)

	err := r.ParseForm()
	if err != nil {
		logger.Error(err.Error())
		return
	}

	codeFromAuth := r.Form.Get("code")
	if codeFromAuth == "" {
		logger.Error("code_from_auth is empty")
		userLoginAuthHandler(w, r)
		return
	}

	formData := url.Values{
		"code":         {codeFromAuth},
		"appid":        {config.OauthAppId},
		"appsecret":    {config.OauthAppSecret},
		"redirect_uri": {config.OauthRedirectUrl},
		"grant_type":   {"auth_code"},
	}
	resp, err := http.PostForm(config.OauthTokenUrl, formData)
	if err != nil {
		logger.Error(err.Error())
		userLoginAuthHandler(w, r)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err.Error())
		userLoginAuthHandler(w, r)
		return
	}
	logger.Info(string(body))

	err = json.Unmarshal(body, &o)
	if err != nil {
		logger.Error(err.Error())
		userLoginAuthHandler(w, r)
		return
	}

	formValue := url.Values {
		"access_token": {o.AccessToken},
		"appid": {config.OauthAppId},
		"openid": {o.OpenId},
	}
	userResp, err := http.PostForm(config.OauthUserUrl, formValue)
	if err != nil {
		logger.Error(err.Error())
		userLoginAuthHandler(w, r)
		return
	}
	userBody, err := ioutil.ReadAll(userResp.Body)
	if err != nil {
		logger.Error(err.Error())
		userLoginAuthHandler(w, r)
		return
	}
	logger.Info(string(userBody))

	var userJson map[string]interface{}
	err = json.Unmarshal(userBody, &userJson)
	if err != nil {
		logger.Error(err.Error())
		userLoginAuthHandler(w, r)
		return
	}

	var authUserName string
	if userName, ok := userJson["name"]; ok {
		if userName != nil {
			authUserName = userName.(string)
		}
	}

	var authUserEmail string
	if userEmail, ok := userJson["email"]; ok {
		if userEmail != nil {
			authUserEmail = userEmail.(string)
		}
	}

	if authUserEmail == "" {
		logger.Error("get email empty")
		userLoginAuthHandler(w, r)
		return
	}

	if !o.checkUserExist(authUserEmail) {
		logger.Error("email not exist in DB")
		http.Redirect(w, r, "/user_login", 302)
		return
	}

	s.Set("is_login", "1")
	s.Set("user_name", authUserName)
	http.Redirect(w, r, "/list_file", 302)
	return
}

func (o *userLoginAuthAPI) checkUserExist(email string) bool {
	db, err := sql.Open("sqlite3", config.SqliteDbPath)
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	defer db.Close()

	querySql := "SELECT id FROM user_list WHERE user_email = ?"
	rows, err := db.Query(querySql, email)
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		count++
	}

	if count > 0 {
		return true
	}

	return false
}