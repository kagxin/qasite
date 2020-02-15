package model

import (
	"github.com/jinzhu/gorm"
)

// Question 问题
type Question struct {
	gorm.Model
	Title  string `gorm:"column:title;type:text;not null"`             // 问题的title
	Text   string `gorm:"column:text; type:text;not null; default ''"` // 问题正文
	UserID uint   `gorm:"column:user_id"`                              // 外键
}

// Comment asdf
type Comment struct {
	gorm.Model
	Text       string `gorm:"column:text; type:text;not null; default ''"`
	QuestionID uint   `gorm:"column:question_id"` //
	UserID     uint   `gorm:"column:user_id"`     //
}
