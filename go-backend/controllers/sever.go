package controllers

import (
	"go-one/go-backend/models"
	"net/http"
)

func StartMainApiServer() error {
	http.HandleFunc("/tasks", models.TaskHandler)
	return http.ListenAndServe(":8888", nil)
}
