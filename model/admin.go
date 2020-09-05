package model

import "gorm.io/gorm"

type Administrators struct {
	gorm.Model
	UserName string `json:"user_name"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}
