package service

import (
	"ginvue/database"
	"ginvue/model"
	"ginvue/serializer"
	"golang.org/x/crypto/bcrypt"
)

//用户登录
type UserLogin struct {
	UserName string `form:"user_name" json:"user_name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

//用户登录的方法
func (s *UserLogin) Login() serializer.Response {
	var u model.Users
	if err := database.Db.Model(&model.Users{}).Where("user_name = ?", s.UserName).First(&u).Error; err != nil {
		return serializer.Response{
			Code: 0,
			Msg:  "该用户不存在",
		}
	}
	//检查密码是否正确
	if bcrypt.CompareHashAndPassword([]byte(u.PassWord), []byte(s.Password)) != nil {
		return serializer.Response{
			Code: 0,
			Msg:  "密码错误",
		}

	}

}
