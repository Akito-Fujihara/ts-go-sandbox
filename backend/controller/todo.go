package controller

import (
	"net/http"

	"github.com/Akito-Fujihara/ts-go-sandbox/database"
	"github.com/Akito-Fujihara/ts-go-sandbox/model"
	"github.com/labstack/echo/v4"
)

func CreateTask(c echo.Context) error {
	todo := model.Todo{}
	if err := c.Bind(&todo); err != nil {
		return err
	}
	database.DB.Create(&todo)
	return c.JSON(http.StatusCreated, todo)
}
