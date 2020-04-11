package databases

import (
	"fmt"
	"log"
	"qasite/conf"

	"github.com/jinzhu/gorm"
)

// InitMysql 初始化gorm db
func InitMysql(mysql *conf.MysqlConfig) *gorm.DB {

	var err error
	var db *gorm.DB

	sqlString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		mysql.Username, mysql.Password, mysql.Host, mysql.Port, mysql.Database)
	fmt.Println(sqlString)
	db, err = gorm.Open("mysql", sqlString)

	if err != nil {
		log.Fatal("mysql connect error %v", err)
	}

	return db
}
