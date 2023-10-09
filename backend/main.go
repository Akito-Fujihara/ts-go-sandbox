package main

import (
	"github.com/Akito-Fujihara/ts-go-sandbox/controller"
	"github.com/Akito-Fujihara/ts-go-sandbox/database"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db, _ := database.DB.DB()
	defer db.Close()

	e.GET("/users", controller.GetUsers)
	e.GET("/users/:id", controller.GetUser)
	e.POST("/users", controller.CreateUser)
	e.PUT("/users/:id", controller.UpdateUser)
	e.DELETE("/users/:id", controller.DeleteUser)
	e.Logger.Fatal(e.Start(":8080"))
}
