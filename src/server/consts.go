package main

import (
	"mtlog"
	"server/session"
	"strconv"
)

// hello world, the web server
var logger *mtlog.Logger
var manager session.Manager
var mtLogPath = "/Users/nature/cluster/centos/osx_work/go/http_server/logs"
var mtLogName = "server"
var mtLogMaxFileSize int64 = 100 * 1024 * 1024
var mtLogKeepFileCount = 100
var httpTemplatePath = "./template"
var httpDataPath = "/Users/nature/cluster/centos/osx_work/go/http_server/data"
var maxUploadSize int64 = 3 * 1024 * 1024 * 1024
var maxUploadSizeStr = strconv.FormatFloat(float64(maxUploadSize) / 1024 / 1024, 'f', 2, 64)
var sqliteDbPath = "/Users/nature/cluster/centos/osx_work/go/http_server/db/data.db"

