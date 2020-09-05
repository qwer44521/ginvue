package router

import (
	"ginvue/controller/admin"
	"ginvue/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRouter(r *gin.Engine) {
	//跨域中间件
	r.Use(middleware.Cors())
	//路由组

	//r.POST("/admin/login", admin.AdminLogin)
	g := r.Group("/admin")
	{
		g.POST("/addadmin", admin.Addadmin)
		g.POST("/login", admin.AdminLogin)

	}

}
