package service

import (
	"fmt"
	"net/http"

	"github.com/delpi/todo-app/internal/user/model"
	"github.com/delpi/todo-app/internal/utils"
)

type IFUserService interface {
	GetUserById(id int) (*model.User, error)
}

type UserService struct {
	Client *http.Client
}

func NewUserService(c *http.Client) *UserService {
	return &UserService{Client: c}
}

func (u *UserService) GetUserById(id int) (*model.User, error) {
	var user model.User
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", id)
	if err := utils.GetJSON(u.Client, url, &user); err != nil {
		return nil, err
	}
	return &user, nil
}
