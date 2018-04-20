package main

import (
	"time"
	"strconv"
)

type tableRow struct{
	Id int `json:"id"`
	FileName string `json:"file_name"`
	UrlName string `json:"url_name"`
	FileUrl string `json:"file_url"`
	Version string `json:"version"`
	Md5 string `json:"md5_value"`
	UserName string `json:"user_name"`
	CreateTimeFmt string `json:"create_time"`
	UpdateTimeFmt string `json:"update_time"`
	Desc string `json:"all_desc"`

	*pageData
	createTime string
	updateTime string
}

func (o *tableRow) format() {
	o.FileUrl = "/data/" + o.UrlName

	createTime, err := strconv.ParseInt(o.createTime, 10, 64)
	if err != nil {
		return
	}

	when := time.Unix(createTime, 0)
	o.CreateTimeFmt = when.Format("2006-01-02 15:04:05")

	updateTime, err := strconv.ParseInt(o.updateTime, 10, 64)
	if err != nil {
		return
	}

	when = time.Unix(updateTime, 0)
	o.UpdateTimeFmt = when.Format("2006-01-02 15:04:05")
}


type jsonListFileAPI struct {
	Id int `json:"id"`
	FileName string `json:"file_name"`
	FileUrl string `json:"file_url"`
	Version string `json:"version"`
	Md5 string `json:"md5_value"`
	CreateTimeFmt string `json:"create_time"`

	urlName string
	createTime string
}

func (o *jsonListFileAPI) format() {
	o.FileUrl = "/data/" + o.urlName

	createTime, err := strconv.ParseInt(o.createTime, 10, 64)
	if err != nil {
		return
	}

	when := time.Unix(createTime, 0)
	o.CreateTimeFmt = when.Format("2006-01-02 15:04:05")
}