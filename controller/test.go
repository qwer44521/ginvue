package controller

import (
	"ginvue/database"
	"ginvue/model"
)

func Test() {
	database.Db.Create(model.User{})
}
