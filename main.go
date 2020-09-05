package main

import (
	"ginvue/database"
	"ginvue/router"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	r := router.Init()
	gin.SetMode(gin.DebugMode)
	_ = r.Run(":8090")
}
