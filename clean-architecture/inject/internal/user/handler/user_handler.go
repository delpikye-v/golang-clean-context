package handler

import (
	"fmt"

	"github.com/delpi/todo-app/internal/user/service"
)

type UserHandler struct {
	Service service.IFUserService
}

func NewUserHandler(s service.IFUserService) *UserHandler {
	return &UserHandler{Service: s}
}

func (h *UserHandler) PrintUser(id int) {
	user, err := h.Service.GetUserById(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("User: %d - %s - %s\n", user.ID, user.Name, user.Email)
}
