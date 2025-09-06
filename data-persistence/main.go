package main

import (
	"github.com/delpi/todo-app/internal/infra"
	"github.com/delpi/todo-app/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	infra.Connect()
	r := gin.Default()

	r.GET("/users", service.GetUser)
	r.POST("/users", service.CreateUser)

	r.GET("/todos", service.GetTodos)
	r.POST("/todos", service.CreateTodo)

	r.Run(":8080")
}
