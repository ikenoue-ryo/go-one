package main

import (
	"fmt"
	"go-one/app/models"
)

func main() {
	fmt.Println(models.Db)
	// controllers.StartMainServer()

	u := &models.User{}
	u.Name = "test1"
	u.Email = "test1@gmail.com"
	u.PassWord = "test1"
	fmt.Println(u)

	u.CreateUser()
	fmt.Println(u)

}
