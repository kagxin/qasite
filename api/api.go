package api

import (
	"qasite/conf"
	"qasite/databases"
)

// Service 接口的handle
type Service struct {
	Conf  *conf.Config
	Mysql *databases.DB
}

// Init Service
func New(conf *conf.Config) *Service {

	return &Service{Conf: conf, Mysql: databases.New(conf.Mysql)}
}
