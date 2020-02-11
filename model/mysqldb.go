package model

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DB * gorm.DB
var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/qasite?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}
}
