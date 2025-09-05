package service

import (
	"fmt"
	"net/http"

	"github.com/delpi/todo-app/internal/todo/model"
	"github.com/delpi/todo-app/internal/user/service"
	"github.com/delpi/todo-app/internal/utils"
)

type TodoService struct {
	Client  *http.Client
	UserSvc service.IFUserService
}

func NewTodoService(c *http.Client, u service.IFUserService) *TodoService {
	return &TodoService{
		Client:  c,
		UserSvc: u,
	}
}

func (t *TodoService) GetTodosByUser(userID int) ([]model.Todo, error) {
	if _, err := t.UserSvc.GetUserById(userID); err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}

	var todos []model.Todo
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/todos?userId=%d", userID)
	if err := utils.GetJSON(t.Client, url, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}
