package model

import (
	"time"
)

//User 用户表
type UserRsp struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
	Username  string    `json:"username"`
}
