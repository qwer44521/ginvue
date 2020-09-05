package controller

import (
	"fmt"
	"ginvue/database"
	"ginvue/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Res struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func returnMsg(ctx *gin.Context, code int, data interface{}, msg string) {
	ctx.JSON(200, Res{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Test202(ctx *gin.Context) {
	ctx.JSON(200, Res{
		Code: 200,
		Data: nil,
		Msg:  "",
	})
}

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
