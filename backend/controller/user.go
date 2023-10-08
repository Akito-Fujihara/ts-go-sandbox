package controller

import (
	"net/http"

	"github.com/Akito-Fujihara/ts-go-sandbox/model"
	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	model.DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}
