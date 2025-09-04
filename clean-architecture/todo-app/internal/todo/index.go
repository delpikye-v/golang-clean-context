package todo

import (
	"github.com/delpi/todo-app/internal/todo/services"
	"github.com/gin-gonic/gin"
)

func RegisterTodoRoutes(r *gin.Engine) {
	t := r.Group("/todos")
	{
		t.GET("", services.GetTodos)
		t.GET("/:id", services.GetTodo)
		t.POST("", services.CreateTodo)
		t.PUT("/:id", services.UpdateTodo)
		t.DELETE("/:id", services.DeleteTodo)
	}
}
