package main

import (
	"fmt"
	"go-one/app/models"
)

func main() {
	fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@gmail.com"
	// u.PassWord = "pass"
	// fmt.Println(u)

	// u.CreateUser()

	u, _ := models.GetUser(1)
	fmt.Println(u)
	// controllers.StartMainServer()

}
