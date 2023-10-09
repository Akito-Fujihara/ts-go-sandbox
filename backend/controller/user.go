package controller

import (
	"errors"
	"net/http"

	"github.com/Akito-Fujihara/ts-go-sandbox/database"
	"github.com/Akito-Fujihara/ts-go-sandbox/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	database.DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}

func GetUsers(c echo.Context) error {
	users := []model.User{}
	database.DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	database.DB.Take(&user)
	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	user := model.User{}
	id := c.Param("id")

	if err := database.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "Record Not found")
		}

		return err
	}
	if err := c.Bind(&user); err != nil {
		return err
	}

	if err := database.DB.Updates(&user).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	user := model.User{}
	id := c.Param("id")

	if err := database.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "Record Not found")
		}

		return err
	}
	if err := database.DB.Delete(&user).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}
