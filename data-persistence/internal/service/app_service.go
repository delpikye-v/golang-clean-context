package service

import (
	"net/http"

	"github.com/delpi/todo-app/internal/infra"
	"github.com/delpi/todo-app/internal/model"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var user []model.User

	infra.DB.Preload("todos").Find(&user)
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	infra.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

// Todos
func GetTodos(c *gin.Context) {
	var todos []model.Todo
	infra.DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var todo model.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	infra.DB.Create(&todo)
	c.JSON(http.StatusCreated, todo)
}
