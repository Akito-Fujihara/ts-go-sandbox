package controller

import (
	"errors"
	"net/http"

	"github.com/Akito-Fujihara/ts-go-sandbox/database"
	"github.com/Akito-Fujihara/ts-go-sandbox/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateTask(c echo.Context) error {
	todo := model.Todo{}
	if err := c.Bind(&todo); err != nil {
		return err
	}
	database.DB.Create(&todo)
	return c.JSON(http.StatusCreated, todo)
}

func GetTasks(c echo.Context) error {
	todos := []model.Todo{}
	database.DB.Find(&todos)
	return c.JSON(http.StatusOK, todos)
}

func GetTask(c echo.Context) error {
	todo := model.Todo{}
	if err := c.Bind(&todo); err != nil {
		return err
	}
	database.DB.Take(&todo)
	return c.JSON(http.StatusOK, todo)
}

func UpdateTask(c echo.Context) error {
	todo := model.Todo{}
	id := c.Param("id")

	if err := database.DB.First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "Record Not found")
		}

		return err
	}
	if err := c.Bind(&todo); err != nil {
		return err
	}

	if err := database.DB.Updates(&todo).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, todo)
}

func DeleteTask(c echo.Context) error {
	todo := model.Todo{}
	id := c.Param("id")

	if err := database.DB.First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "Record Not found")
		}

		return err
	}
	if err := database.DB.Delete(&todo).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, todo)
}
