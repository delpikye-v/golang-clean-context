package main

import (
	db "github.com/delpi/todo-app/db"
	"github.com/delpi/todo-app/internal/todo"
	"github.com/delpi/todo-app/internal/users"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	r := gin.Default()

	todo.RegisterTodoRoutes(r)
	users.RegisterTodoRoutes(r)

	r.Run(":8080")
}
