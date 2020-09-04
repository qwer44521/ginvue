package main

import (
	"ginvue/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	_ = r.Run(":8090")
}
