package api

import (
	"qasite/conf"
	"qasite/databases"
)

// Service 接口的handle
type Service struct {
	conf *conf.Config
	db   *databases.DB
}

// Init Service
func Init(conf *conf.Config, db *databases.DB) *Service {
	return &Service{conf, db}
}
