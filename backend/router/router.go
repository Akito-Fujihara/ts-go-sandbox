package router

import (
	"github.com/Akito-Fujihara/ts-go-sandbox/controller"
	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	// TODO
	e.GET("/todos", controller.GetTasks)
	e.GET("/todos/:id", controller.GetTask)
	e.POST("/todos", controller.CreateTask)
	e.PUT("/todos/:id", controller.UpdateTask)
	e.DELETE("/todos/:id", controller.DeleteTask)

	return e
}
