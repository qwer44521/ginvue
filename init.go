package main

import (
	"ginvue/middleware"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine  {
	r :=gin.Default()
	//跨域中间件
	r.Use(middleware.Cors())
	//路由组
	api := r.Group("api"){
		api.POST()
	}
	return r
}
