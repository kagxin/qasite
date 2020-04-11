package main

import (
	"log"
	"qasite/api"
	"qasite/conf"
	"qasite/databases"
	"qasite/model"
	"qasite/router"
)

func main() {
	config := conf.Init()
	db := databases.InitMysql(config.Mysql)
	model.CreateTable(db)
	srv := api.Init(config, db)
	r := router.Router(srv)
	if err := r.Run(); err != nil {
		log.Fatal("gin run failed.")
	}
}
