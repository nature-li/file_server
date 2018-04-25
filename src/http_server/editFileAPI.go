package main

import (
	"database/sql"
	"session"
	"net/http"
	"encoding/json"
)

type editFileAPI struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	session session.Session
}

func (o *editFileAPI) queryFileList(queryId string) *tableRow {
	db, err := sql.Open("sqlite3", config.SqliteDbPath)
	if err != nil {
		logger.Error(err.Error())
		return nil
	}
	defer db.Close()

	querySql := "SELECT id,file_name,file_size,url_name,version,md5_value,user_name,desc,create_time,update_time FROM file_list WHERE id = ?"
	rows, err := db.Query(querySql, queryId)
	if err != nil {
		logger.Error(err.Error())
		return nil
	}
	defer rows.Close()

	row := &tableRow{}
	for rows.Next() {
		err = rows.Scan(&row.Id, &row.FileName, &row.FileSize, &row.UrlName, &row.Version, &row.Md5, &row.UserName, &row.Desc, &row.createTime, &row.updateTime)
		if err != nil {
			logger.Error(err.Error())
			return nil
		}
		break
	}
	row.format()

	return row
}

func (o *editFileAPI) editFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Info(r.Form.Encode())

	fileId := r.Form.Get("file_id")
	if fileId == "" {
		o.render(w, false, "FILE_ID_EMPTY")
		return
	}
	fileVersion := r.Form.Get("file_version")
	fileDesc := r.Form.Get("file_desc")

	// 检测数据长度
	version := []rune(fileVersion)
	if len(version) > MAX_VERSION_LEN {
		o.render(w, false, "FILE_VERSION_BIG")
		return
	}
	desc := []rune(fileDesc)
	if len(desc) > MAX_DESC_LEN {
		o.render(w, false, "FILE_DESC_BIG")
		return
	}

	db, err := sql.Open("sqlite3", config.SqliteDbPath)
	if err != nil {
		logger.Error(err.Error())
		o.render(w, false, "OPEN_DB_FAILED")
		return
	}
	defer db.Close()

	if o.editDB(db, fileId, fileVersion, fileDesc) {
		o.render(w, true, "SUCCESS")
	} else {
		o.render(w, false, "EDIT_DB_FAILED")
	}
}

func (o *editFileAPI) editDB(db *sql.DB, fileId, fileVersion, fileDesc string) bool  {
	querySql := "update file_list set version=?, desc=? where id=?"
	logger.Info(querySql)
	rows, err := db.Exec(querySql, fileVersion, fileDesc, fileId)
	if err != nil {
		logger.Error(err.Error())
		return false
	}

	count, err := rows.RowsAffected()
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Infof("affected rows: %v", count)
	return true
}

func (o *editFileAPI) render(w http.ResponseWriter, success bool, msg string) {
	o.Success = success
	o.Msg = msg

	result, err := json.Marshal(o)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	w.Write(result)
}
