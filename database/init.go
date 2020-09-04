package database

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var Db *gorm.DB

//连接数据库
func Init() {
	_ = godotenv.Load()
	db, err := gorm.Open(mysql.Open(os.Getenv("DB_DSN")), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	Db = db
}
