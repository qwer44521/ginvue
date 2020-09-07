package main

import (
	"ginvue/database"
	"ginvue/router"
	"github.com/gin-gonic/gin"
)

//@title 测试数据12
func main() {
	//controller.Testtoken()
	database.Init()
	r := router.Init()
	gin.SetMode(gin.DebugMode)
	_ = r.Run(":8090")
}
