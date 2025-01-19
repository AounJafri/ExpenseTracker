package config

import (
	"database/sql"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {

	d, err := sql.Open("mysql", "root:aoun123@tcp(127.0.0.1:3306)/ExpenseTracker")

	if err != nil {
		fmt.Println(err.Error())
	}

	return d
}
