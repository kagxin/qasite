package model

import (
	"github.com/jinzhu/gorm"
)

//User 用户表
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique_index;not null"` // 列名为 `username`
	Password string `json:"password" gorm:"not null" `             // 列名为 `password`
}
