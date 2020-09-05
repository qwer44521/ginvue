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

//添加管理员
type Administrators struct {
	UserName string `form:"user_name" json:"user_name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Nickname string `form:"nickname" json:"nickname"`
}
type Adminlogin struct {
	UserName string `form:"user_name" json:"user_name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

//添加管理员方法
func (s *Administrators) Addadmin() serializer.Response {
	u := model.Administrators{
		UserName: s.UserName,
		Password: s.Password,
		Nickname: s.Nickname,
	}
	//密码加密
	bytes, err := bcrypt.GenerateFromPassword([]byte(s.Password), 12)
	if err != nil {
		return serializer.Response{
			Code: 0,
			Msg:  "密码加密失败",
		}
	}
	u.Password = string(bytes)
	//创建管理员
	if err := database.Db.Model(&model.Administrators{}).Create(&u).Error; err != nil {
		return serializer.Response{
			Code: 0,
			Msg:  "注册失败",
		}
	}
	return serializer.Response{
		Code: 1,
		Msg:  "注册成功",
	}
}

//管理员登录方法
func (s *Adminlogin) Adminlogin() serializer.Response {
	var u model.Administrators
	var j middleware.JWT
	//检查管理员是否存在
	if err := database.Db.Model(&model.Administrators{}).Where("user_name = ?", s.UserName).First(&u).Error; err != nil {
		return serializer.Response{
			Code: 0,
			Msg:  "管理员不存在",
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
			Msg:  "登陆失败：token未生成",
		}

	}
	return serializer.Response{
		Code: 1,
		Msg:  "登陆成功",
		Data: token,
	}
}
