package main

import (
	"net/http"
	"encoding/json"
	"database/sql"
	_"os"
	_"path/filepath"
	"path/filepath"
	"os"
	"server/session"
)

type deleteFileAPI struct {
	session session.Session
	Success bool `json:"success"`
	Desc string `json:"desc"`
}

func (o *deleteFileAPI) handle(w http.ResponseWriter, r *http.Request) {
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

	db, err := sql.Open("sqlite3", config.SqliteDbPath)
	if err != nil {
		logger.Error(err.Error())
		o.render(w, false, "OPEN_DB_FAILED")
		return
	}
	defer db.Close()

	if o.deleteFromDisk(w, db, fileId) {
		o.deleteFromDB(w, db, fileId)
	}
}

func (o *deleteFileAPI) render(w http.ResponseWriter, success bool, desc string) {
	o.Success = success
	o.Desc = desc

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


func (o *deleteFileAPI) deleteFromDisk(w http.ResponseWriter, db *sql.DB, fileId string) bool {
	querySql := "SELECT url_name FROM file_list WHERE id = ?"
	rows, err := db.Query(querySql, fileId)
	if err != nil {
		logger.Error(err.Error())
		o.render(w, false, "QUERY_DB_FAILED")
		return false
	}
	defer rows.Close()

	var urlName string
	for rows.Next() {
		err = rows.Scan(&urlName)
		if err != nil {
			logger.Error(err.Error())
			o.render(w, false, "SCAN_DB_FAILED")
			return false
		}
		break
	}

	if urlName == "" {
		logger.Error("fileId" + fileId + " not exist")
		o.render(w, false, "FILE_NOT_EXISTS")
		return false
	}

	fullPath := filepath.Join(config.UploadDataPath, urlName)
	logger.Info("remove file: " + fullPath)
	err = os.Remove(fullPath)
	if err != nil {
		logger.Warn(err.Error())
	}

	return true
}

func (o *deleteFileAPI) deleteFromDB(w http.ResponseWriter, db *sql.DB, fileId string) bool {
	querySql := "DELETE FROM file_list WHERE id = ?"
	stmt, err := db.Prepare(querySql)
	if err != nil {
		logger.Error(err.Error())
		o.render(w, false, "DB_PREPARE_FAILED")
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(fileId)
	if err != nil {
		logger.Error(err.Error())
		o.render(w, false, "DB_EXEC_FAILED")
	}

	o.render(w, true, "SUCCESS")
	return true
}