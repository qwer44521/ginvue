package service

import (
	"ginvue/database"
	"ginvue/middleware"
	"ginvue/model"
	"ginvue/serializer"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

//用户登录
type UserLogin struct {
	UserName string `form:"user_name" json:"user_name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

//用户登录的方法
func (s *UserLogin) Login() serializer.Response {
	var u model.Users
	var j middleware.JWT
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
	//生成token
	token, err := j.CreateToken(middleware.CustomClaims{
		ID:       u.ID,
		UserName: u.UserName,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),
			ExpiresAt: int64(time.Now().Unix() + 3600),
			Issuer:    os.Getenv("JWT_ISSUER"),
		},
	})
	if err != nil {
		return serializer.Response{
			Code: -1,
			Msg:  "登陆失败，未生成token",
		}
	}
	return serializer.Response{
		Code: 1,
		Msg:  "登录成功",
		Data: token,
	}
}
