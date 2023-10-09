package database

import (
	"log"

	"github.com/Akito-Fujihara/ts-go-sandbox/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	dsn := "root:password@tcp(db:3306)/maindb?charset=utf8mb4&parseTime=True"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(dsn + " database can't connect")
	}

	user := model.User{}
	DB.AutoMigrate(&user)
}
