package main

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
)

type authResult struct {
	AccessToken string `json:"access_token"`
	OpenId      string `json:"openid"`
}

func userLoginAuthAPIHandler(w http.ResponseWriter, r *http.Request) {
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
		"appid":        {serverAuthAppId},
		"appsecret":    {serverAuthAppSecret},
		"redirect_uri": {serverAuthRedirectUrl},
		"grant_type":   {"auth_code"},
	}
	resp, err := http.PostForm(serverAuthTokenUrl, formData)
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

	var result authResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		logger.Error(err.Error())
		userLoginAuthHandler(w, r)
		return
	}

	formValue := url.Values {
		"access_token": {result.AccessToken},
		"appid": {serverAuthAppId},
		"openid": {result.OpenId},
	}
	userResp, err := http.PostForm(serverAuthUserUrl, formValue)
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

	if userName, ok := userJson["name"]; ok {
		if userName != nil {
			s.Set("is_login", "1")
			s.Set("user_name", userName.(string))
			http.Redirect(w, r, "/user_login", 302)
			return
		}
	}

	logger.Error(err.Error())
	userLoginAuthHandler(w, r)
	return
}
