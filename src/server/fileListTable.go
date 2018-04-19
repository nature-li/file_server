package main

import (
	"time"
	"strconv"
)

type tableRow struct {
	Id int `json:"id"`
	FileName string `json:"file_name"`
	UrlName string `json:"file_url"`
	Version string `json:"version"`
	Md5 string `json:"md5_value"`
	UserName string `json:"user_name"`
	CreateTime string `json:"create_time_secs"`
	CreateTimeFmt string `json:"create_time"`
	UpdateTime string `json:"update_time_secs"`
	UpdateTimeFmt string `json:"update_time"`
	Desc string `json:"all_desc"`
	ShortDesc string `json:"desc"`
}

func (o *tableRow) format() {
	if len(o.Desc) > 30 {
		o.ShortDesc = o.Desc[0:30]
	} else {
		o.ShortDesc = o.Desc
	}

	createTime, err := strconv.ParseInt(o.CreateTime, 10, 64)
	if err != nil {
		return
	}

	when := time.Unix(createTime, 0)
	o.CreateTimeFmt = when.Format("2006-01-02 15:04:05")

	updateTime, err := strconv.ParseInt(o.UpdateTime, 10, 64)
	if err != nil {
		return
	}

	when = time.Unix(updateTime, 0)
	o.UpdateTimeFmt = when.Format("2006-01-02 15:04:05")
}