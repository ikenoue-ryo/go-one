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

	// controllers.StartMainServer()
	// user, _ := models.GetUser(2)
	// user.CreateTodo("Second todo")

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// user, _ := models.GetUser(2)
	// todos, _ := user.GetTodosByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// t, _ := models.GetTodo(1)
	// t.Content = "Update Todo"
	// t.UpdateTodo()

	t, _ := models.GetTodo(2)
	t.DeleteTodo()
}
