package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"strconv"
)

type dataItem struct {
	Id int `json:"id"`
	FileName string `json:"file_name"`
	FileUrl string `json:"file_url"`
	Version string `json:"version"`
	Md5 string `json:"md5"`
	UserName string `json:"user_name"`
	CreateTime string `json:"create_time"`
	Desc string `json:"desc"`
}

type dataResult struct {
	Success string `json:"success"`
	ItemCount int `json:"item_count"`
	Content []dataItem `json:"content"`
}

func queryFileListHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fileName := r.Form.Get("file_name")
	offSet := r.Form.Get("off_set")
	limit := r.Form.Get("limit")

	fmt.Println(fileName)
	fmt.Println(offSet)
	fmt.Println(limit)

	count, _ := strconv.Atoi(limit)
	result := &dataResult{Success:"true", ItemCount:100}
	for i := 0; i < count; i++ {
		it := dataItem{}
		it.Id = i
		it.FileName = "hello.txt"
		it.FileUrl = "/data/1524059667_chart.png"
		it.Version = "0.12"
		it.Md5 = "md5"
		it.UserName = "lyg"
		it.CreateTime = "2018"
		it.Desc = "desc"
		result.Content = append(result.Content, it)
	}

	jsonString, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, string(jsonString))
}
