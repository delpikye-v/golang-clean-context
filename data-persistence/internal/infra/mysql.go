package infra

import (
	"fmt"

	"github.com/delpi/todo-app/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	dsn := "root:root1234@tcp(127.0.0.1:3306)/todo_test"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&model.User{}, &model.Todo{})

	fmt.Println("✅ Connected to MySQL")
}
