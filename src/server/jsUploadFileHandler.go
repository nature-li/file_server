package main

import (
	"net/http"
	"path/filepath"
	"os"
	"encoding/json"
	"time"
	"strconv"
	"io"
	"fmt"
)

const maxUploadSize = 3 * 1024 * 1024 * 1024
const uploadPath = "/Users/liyanguo/code/GoglandProjects/http_server/data"

type jsUploadFileResult struct {
	RetCode int `json:"code"`
	Desc string `json:"desc"`
}

func jsUploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// 检测文件大小
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
		return
	}

	// 打开旧文件
	srcFile, handler, err := r.FormFile("uploadFile")
	if err != nil {
		renderError(w, "INVALID_FILE", http.StatusBadRequest)
		return
	}
	defer srcFile.Close()

	// 打开新文件
	fileName := expendFileName(handler.Filename)
	newPath := filepath.Join(uploadPath, fileName)
	dstFile, err := os.Create(newPath)
	if err != nil {
		renderError(w, "CREATE_FILE_ERROR", http.StatusInternalServerError)
		return
	}
	defer dstFile.Close()

	// 复制文件
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		renderError(w, "SAVE_FILE_ERROR", http.StatusBadRequest)
		return
	}

	// 返回成功
	renderSuccess(w, "SUCCESS")
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)

	result := jsUploadFileResult{RetCode:statusCode, Desc:message}
	byteResult, _ := json.Marshal(result)
	w.Write(byteResult)
}

func renderSuccess(w http.ResponseWriter, message string) {
	result := jsUploadFileResult{RetCode:0, Desc:message}
	byteResult, _ := json.Marshal(result)
	w.Write(byteResult)
}

func expendFileName(fileName string) string {
	now := time.Now().Unix()
	nowStr := strconv.FormatInt(now, 10)
	return nowStr + "_" + fileName
}

