package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"user_name"`
	PassWord string `json:"password"`
	Time     string `json:"time"`
	Status   string `json:"status"`
}
