package main

import (
	model "qasite/model"
	"qasite/router"
)

func init() {
	model.DB.SingularTable(true)
	model.DB.AutoMigrate(&model.User{}, &model.Question{}, &model.Comment{})
}

func main() {

	defer model.DB.Close()
	model.DB.LogMode(true)
	r := router.Router()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

/*
TODO:
1、response整理 {code:200, message:"", data:{}}
2、response 字段名大写to小写
3、gorm 多表查询
4、不规则json解析
5、gorm result row
*/
