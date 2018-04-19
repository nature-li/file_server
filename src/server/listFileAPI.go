package main

import (
	"net/http"
	"encoding/json"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type listFileAPI struct {
	Success   string     `json:"success"`
	ItemCount int        `json:"item_count"`
	Content   []tableRow `json:"content"`
}

func (o *listFileAPI) handle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fileName := r.Form.Get("file_name")
	offset := r.Form.Get("off_set")
	limit := r.Form.Get("limit")

	logger.Info(r.Form.Encode())

	db, err := sql.Open("sqlite3", sqliteDbPath)
	if err != nil {
		logger.Error(err.Error())
		o.render(w, "false", 0, nil)
	}
	defer db.Close()

	code, totalCount := o.countDB(db, fileName)
	if code != http.StatusOK {
		o.render(w, "false", 0, nil)
		return
	}

	code, _, rows := o.queryDB(db, fileName, limit, offset)
	if code != http.StatusOK {
		o.render(w, "false", 0, nil)
		return
	}

	o.render(w, "true", totalCount, rows)
}

func (o *listFileAPI)render(w http.ResponseWriter, success string, itemCount int, content []tableRow) {
	w.WriteHeader(http.StatusOK)

	o.Success = success
	o.ItemCount = itemCount
	o.Content = content
	byteResult, err := json.Marshal(o)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	w.Write(byteResult)
}

func (o *listFileAPI) countDB(db *sql.DB, fileName string) (code int, totalCount int) {
	var err error
	var querySql string
	var rows *sql.Rows
	if len(fileName) != 0 {
		querySql = "SELECT COUNT(1) AS COUNT FROM file_list WHERE file_name like ?"
		logger.Info(querySql)
		rows, err = db.Query(querySql, "%" + fileName + "%")
	} else {
		querySql = "SELECT COUNT(1) AS COUNT FROM file_list"
		logger.Info(querySql)
		rows, err = db.Query(querySql, fileName)
	}
	if err != nil {
		logger.Error(err.Error())
		return http.StatusInternalServerError, 0
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&totalCount)
	}

	return http.StatusOK, totalCount
}

func (o *listFileAPI) queryDB(db *sql.DB, fileName, limit, offset string) (int, string, []tableRow) {
	var err error
	var dataSql string
	if len(fileName) != 0 {
		dataSql = "SELECT id,file_name,url_name,version,md5_value,user_name,desc,create_time,update_time " +
			"FROM file_list WHERE file_name like ? order by create_time limit ? offset ?"
	} else {
		dataSql = "SELECT id,file_name,url_name,version,md5_value,user_name,desc,create_time,update_time " +
			"FROM file_list order by create_time limit ? offset ?"
	}

	var rows *sql.Rows
	if len(fileName) != 0 {
		logger.Info(dataSql)
		rows, err = db.Query(dataSql, "%" + fileName + "%", limit, offset)
	} else {
		logger.Info(dataSql)
		rows, err = db.Query(dataSql, limit, offset)
	}
	if err != nil {
		logger.Error(err.Error())
		return http.StatusInternalServerError, "QUERY_DB_FAILED", nil
	}
	defer rows.Close()

	var rowList []tableRow
	for rows.Next() {
		row := tableRow{}
		err = rows.Scan(&row.Id, &row.FileName, &row.UrlName, &row.Version, &row.Md5, &row.UserName, &row.Desc, &row.CreateTime, &row.UpdateTime)
		if err != nil {
			logger.Error(err.Error())
			return http.StatusInternalServerError, "QUERY_DB_FAILED", nil
		}

		row.format()
		rowList = append(rowList, row)
	}

	return http.StatusOK, "SUCCESS", rowList
}
