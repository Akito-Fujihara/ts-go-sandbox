package router

import (
	"github.com/Akito-Fujihara/ts-go-sandbox/controller"
	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.GET("/users", controller.GetUsers)
	e.GET("/users/:id", controller.GetUser)
	e.POST("/users", controller.CreateUser)
	e.PUT("/users/:id", controller.UpdateUser)
	e.DELETE("/users/:id", controller.DeleteUser)

	return e
}
