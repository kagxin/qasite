package api

import (
	"github.com/jinzhu/gorm"
	"qasite/conf"
)

// Service 接口的handle
type Service struct {
	Conf *conf.Config
	DB   *gorm.DB
}

// Init Service
func Init(conf *conf.Config, db *gorm.DB) *Service {
	return &Service{conf, db}
}
