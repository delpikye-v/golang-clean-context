package handler

import (
	"fmt"

	"github.com/delpi/todo-app/internal/todo/service"
)

type TodoHandler struct {
	Service *service.TodoService
}

func NewTodoHandler(s *service.TodoService) *TodoHandler {
	return &TodoHandler{Service: s}
}

func (h *TodoHandler) PrintTodos(userID int) {
	todos, err := h.Service.GetTodosByUser(userID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Todos for user %d:\n", userID)
	for _, t := range todos {
		fmt.Printf("- %s (completed: %v)\n", t.Title, t.Completed)
	}
}
