package main

import (
	"go-one/go-backend/database"
	"go-one/go-backend/models"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	database.Connect()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	e.GET("/category", models.GetCategory)
	e.GET("/users", models.GetUsers)
	e.POST("/users", models.CreateUser)

	e.Logger.Fatal(e.Start(":1323"))
}
