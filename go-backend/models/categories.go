package models

import (
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

func GetCategory(c echo.Context) error {
	return c.JSON(http.StatusOK, categories)
}
