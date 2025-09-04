package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// for test => move to service
func index(c *gin.Context) {
	c.JSON(http.StatusOK, []any{
		map[string]any{
			"name": "abcd",
		},
	})
}

func RegisterTodoRoutes(r *gin.Engine) {
	t := r.Group("/users")
	{
		t.GET("", index)
		// t.GET("/:id", services.GetUser)
	}
}
