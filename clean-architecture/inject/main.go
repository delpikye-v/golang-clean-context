package main

import (
	todoHandler "github.com/delpi/todo-app/internal/todo/handler"
	todoService "github.com/delpi/todo-app/internal/todo/service"
	userHandler "github.com/delpi/todo-app/internal/user/handler"
	userService "github.com/delpi/todo-app/internal/user/service"
	appUtils "github.com/delpi/todo-app/internal/utils"
)

func main() {
	client := appUtils.NewHTTPClient()

	// User
	userSvc := userService.NewUserService(client)
	uHandler := userHandler.NewUserHandler(userSvc)

	// Todo
	todoSvc := todoService.NewTodoService(client, userSvc) // inject IFUserService
	tHandler := todoHandler.NewTodoHandler(todoSvc)

	// RUN
	uHandler.PrintUser(1)
	tHandler.PrintTodos(1)
}
