package model

import (
	"time"
)

//User 用户
type User struct {
	Id       uint64    `json:"id"`
	Username string    `json:"username" xorm:"unique"`
	Nickname string    `json:"nickname,omitempty"`
	Email    string    `json:"email,omitempty"`
	Disabled bool      `json:"disabled"`
	Created  time.Time `json:"created" xorm:"created"`
}

//Password 密码
type Password struct {
	Id       uint64 `json:"id" xorm:"pk"`
	Password string `json:"password"`
}
