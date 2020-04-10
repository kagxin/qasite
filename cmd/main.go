package main

import (
	"qasite/api"
	"qasite/conf"
	"qasite/databases"
	"qasite/router"
)

func main() {
	conf := conf.Init()
	db := databases.InitMysql(conf.Mysql)
	srv := api.Init(conf, db)

	router := router.Router(srv)

	router.Run()
}
