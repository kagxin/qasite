package model

import (
	"github.com/jinzhu/gorm"
)

// Question 问题
type Question struct {
	gorm.Model
	Title  string `gorm:"column:title;type:text;not null"`             // 问题的title
	Text   string `gorm:"column:text; type:text;not null; default ''"` //
	UserID uint   // 外键
}

// Comment asdf
type Comment struct {
	gorm.Model
	Text       string `gorm:"colume:text; type:text; not null; default ''"`
	QuestionID uint
}
