package main

import (
	"database/sql"
)

type editFileAPI struct {
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
