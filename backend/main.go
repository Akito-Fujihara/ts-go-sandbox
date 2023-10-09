package main

import (
	"github.com/Akito-Fujihara/ts-go-sandbox/controller"
	"github.com/Akito-Fujihara/ts-go-sandbox/model"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db, _ := model.DB.DB()
	defer db.Close()

	e.GET("/users", controller.GetUsers)
	e.GET("/users/:id", controller.GetUser)
	e.POST("/users", controller.CreateUser)
	e.Logger.Fatal(e.Start(":8080"))
}
