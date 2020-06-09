package main

import (
	"log"
	"qasite/api"
	"qasite/conf"
	"qasite/model"
	"qasite/router"
)

func main() {
	config := conf.New()
	srv := api.New(config)
	model.CreateTable(srv.Mysql.DB)
	r := router.Router(srv)
	if err := r.Run(); err != nil {
		log.Fatal("gin run failed.")
	}
}
