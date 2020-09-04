package main

import (
	"ginvue/controller"
	"ginvue/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	//controller.Test()
	//controller.Test1()
	controller.Test3()
	r := gin.Default()
	gin.SetMode(gin.DebugMode)

	//r.GET("/test", func(context *gin.Context) {
	//	res := controller.Test1()
	//	context.JSON(http.StatusOK, res)
	//})

	_ = r.Run(":8090")
}
