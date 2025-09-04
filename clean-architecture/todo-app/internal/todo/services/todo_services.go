/*
	This source code is for teaching purposes and demonstrates clean code practices.

+ gorm.io/driver/mysql
+ gorm.io/gorm
*/
package services

import (
	"net/http"
	"strconv"

	db "github.com/delpi/todo-app/db"
	model "github.com/delpi/todo-app/internal/todo/models"
	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	rows, _ := db.DB.Query("SELECT id, task, done FROM todos")
	defer rows.Close()

	var todos []model.Todo
	for rows.Next() {
		var t model.Todo
		rows.Scan(&t.ID, &t.Task, &t.Done)
		todos = append(todos, t)
	}

	c.JSON(http.StatusOK, todos)
}

func GetTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	row := db.DB.QueryRow("SELECT id, task, done FROM todos WHERE id=?", id)

	var t model.Todo
	err := row.Scan(&t.ID, &t.Task, &t.Done)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "model.Todo not found"})
		return
	}

	c.JSON(http.StatusOK, t)
}

func CreateTodo(c *gin.Context) {
	var t model.Todo
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, _ := db.DB.Exec("INSERT INTO todos (task, done) VALUES (?, ?)", t.Task, t.Done)
	id, _ := res.LastInsertId()
	t.ID = int(id)

	c.JSON(http.StatusCreated, t)
}

func UpdateTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var t model.Todo
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Exec("UPDATE todos SET task=?, done=? WHERE id=?", t.Task, t.Done, id)
	t.ID = id

	c.JSON(http.StatusOK, t)
}

func DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	db.DB.Exec("DELETE FROM todos WHERE id=?", id)
	c.Status(http.StatusNoContent)
}
