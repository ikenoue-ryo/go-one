package main

import (
	"fmt"
	"go-one/app/controllers"
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

	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// u.Name = "test2"
	// u.Email = "test2@gmail.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u, _ := models.GetUser(1)
	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	controllers.StartMainServer()
}
