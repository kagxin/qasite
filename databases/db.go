package databases

import (
	"fmt"
	"log"
	"qasite/conf"

	"github.com/jinzhu/gorm"
)

type DB struct {
	DB *gorm.DB
}

// InitMysql 初始化gorm db
func New(mysql *conf.MysqlConfig) *DB {

	var err error
	var db *gorm.DB

	sqlString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		mysql.Username, mysql.Password, mysql.Host, mysql.Port, mysql.Database)
	fmt.Println(sqlString)
	db, err = gorm.Open("mysql", sqlString)

	if err != nil {
		log.Fatalf("mysql connect error %v", err)
	}

	return &DB{DB: db}
}
