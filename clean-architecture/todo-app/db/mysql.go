// This source code is for teaching purposes and demonstrates clean code practices.
package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error
	dsn := "root:root1234@tcp(127.0.0.1:3306)/todo_test"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("✅ Connected to MySQL")
}
