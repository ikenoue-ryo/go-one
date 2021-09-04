package main

import (
	"go-one/go-backend/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Category struct {
	Id   int    `json:"id" param:"id"`
	Name string `json:"name"`
	Kana string `json:"kana"`
}

var categories = []Category{
	{
		Id:   1,
		Name: "椅子",
		Kana: "イス",
	},
	{
		Id:   2,
		Name: "机",
		Kana: "ツクエ",
	},
	{
		Id:   3,
		Name: "ベッド",
		Kana: "ベッド",
	},
}

type User struct {
	Id    int    `json:"id" param:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func getCategory(c echo.Context) error {
	return c.JSON(http.StatusOK, categories)
}

func getUsers(c echo.Context) error {
	users := []User{}
	database.DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func createUser(c echo.Context) error {
	user := User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	database.DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}

func main() {
	e := echo.New()
	database.Connect()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	e.GET("/category", getCategory)
	e.GET("/users", getUsers)
	e.POST("/users", createUser)

	e.Logger.Fatal(e.Start(":1323"))
}
