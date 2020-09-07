package service

import (
	"fmt"
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
	UserName string `form:"user_name" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Nickname string `form:"nickname" json:"nickname"`
}
type Adminlogin struct {
	UserName string `form:"user_name" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
type AdminInfo struct {
	Token string `form:"token" json:"token" binding:"required"`
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
			Msg:  "添加失败",
		}
	}
	return serializer.Response{
		Code: 1,
		Msg:  "管理员添加成功",
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
	//验证密码是否准确
	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(s.Password)) != nil {
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
			Msg:  "登陆失败：token未生成",
		}

	}
	return serializer.Response{
		Code:  1,
		Msg:   "登陆成功",
		Token: token,
	}
}

//获取用户的个人信息（包括用户名、昵称、角色名等等）
func (s AdminInfo) AdminInfo() serializer.Response {
	var u model.Administrators
	var j middleware.JWT

	//解析token
	t, err := j.ParseToken(s.Token)
	if err != nil {
		panic("解析token失败")

	}
	if err := database.Db.Model(&model.Administrators{}).Where("id = ?", t.ID).First(&u).Error; err != nil {
		return serializer.Response{
			Code: 0,
			Msg:  "数据返回失败",
		}

	}
	fmt.Println(u)
	return serializer.Response{
		Code: 1,
		Msg:  "接收数据成功",
		Data: u,
	}
}
