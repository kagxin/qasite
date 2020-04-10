package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

/*
表结文件
*/

//User 用户表
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"` // 列名为 `username`
	Password string `gorm:"not null" `       // 列名为 `password`
}

// Token 用户token
type Token struct {
	gorm.Model
	UserID    uint      `gorm:"column:user_id"`
	Token     string    `gorm:"column:text; type:text;not null;"`
	ExpiresIn time.Time `gorm:"index; not null"`
}

// Question orm 表结构
type Question struct {
	gorm.Model
	Title  string `gorm:"column:title;type:text;not null; default ''"` // 问题的title
	Text   string `gorm:"column:text; type:text;not null; default ''"` // 问题正文
	UserID uint   `gorm:"column:user_id"`                              // 外键
}

// Comment orm 表结构
type Comment struct {
	gorm.Model
	Text       string `gorm:"column:text; type:text;not null; default ''"`
	QuestionID uint   `gorm:"column:question_id"` //
	UserID     uint   `gorm:"column:user_id"`     //
}

// CreateTable 创建表结构
func CreateTable(db *gorm.DB) {
	// 表名为结构体名的小写
	db.SingularTable(true)
	// Migrate the schema
	db.AutoMigrate(&User{}, &Token{}, &Question{}, &Comment{})
}
