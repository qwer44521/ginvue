package controller

import (
	"fmt"
	"ginvue/database"
	"ginvue/middleware"
	"ginvue/model"
	"golang.org/x/crypto/bcrypt"
)

func Test() {
	result := database.Db.Create(model.Test{
		Name: string("zjj14"),
		Pwd:  string("12fdf"),
	})
	println(result.RowsAffected)
}
func Test1() []model.Test {
	var res []model.Test
	if err := database.Db.Find(&res).Error; err != nil {
		panic(err.Error())
	}

	return res

}
func Test2() {
	passwordOK := "admin"
	hash, err := bcrypt.GenerateFromPassword([]byte(passwordOK), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	encodepw := string(hash)
	fmt.Println(encodepw)
}
func Test3() {
	pw := "$2a$10$47jehSCKjwiDQV8qVbHim.bMFps0VHFRVOj9RZfS07pGzhPxc0KPi"
	pp := "admin"
	err := bcrypt.CompareHashAndPassword([]byte(pw), []byte(pp))
	if err != nil {
		fmt.Println("pw错了")
	} else {
		fmt.Println("ok")
	}
}
func Testtoken() {
	var j middleware.JWT
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VyX25hbWUiOiJhZG1pbiIsImV4cCI6MTU5OTM5MDAwNywiaXNzIjoiZ29naW4iLCJuYmYiOjE1OTkzODU0MDd9.MfHNQtovzy0pUu3qW3e-daSLGJ8Rw6-k_KmwXm2aDtQ"
	y, err := j.ParseToken(token)
	if err != nil {
		panic("错误信息")
	}
	fmt.Println(y.ID)
	fmt.Printf("%#v", y)

}
