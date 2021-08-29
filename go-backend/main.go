package main

import (
	"go-one/go-backend/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

// import (
// 	"go-one/go-backend/controllers"
// )

// func main() {
// 	controllers.StartMainApiServer()
// }

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

func getCategory(c echo.Context) error {
	return c.JSON(http.StatusOK, categories)
}

func main() {
	e := echo.New()
	database.Connect()
	sqlDB, _ := database.DB.DB()
	defer sqlDB.Close()

	e.GET("/category", getCategory)

	e.Logger.Fatal(e.Start(":1323"))
}
