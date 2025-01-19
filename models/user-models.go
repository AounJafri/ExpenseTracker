package models

import (
	"database/sql"
	"fmt"

	"github.com/Aounjafri/ExpenseTracker/config"
)

var db *sql.DB

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func init() {
	db = config.Connect()
}

// USER FUNCTIONS

func CreateUser(username, password, email string) error {
	_, err := db.Exec("INSERT INTO User(username,password,email) VALUES(?,?,?)", username, password, email)

	if err != nil {
		fmt.Println(err.Error())
	}

	return err
}

func GetHashedPassAndId(name string) (int, string) {
	var hash string
	var id int
	db.QueryRow("Select id,password from User Where username=?", name).Scan(&id, &hash)
	// fmt.Println(hash)
	return id, hash
}

func GetUsers() []User {
	result, err := db.Query("Select * from User")

	if err != nil {
		fmt.Println(err.Error())
	}
	var users []User

	for result.Next() {
		var use User

		result.Scan(&use.Id, &use.Name, &use.Password, &use.Email)
		// fmt.Println(use)
		users = append(users, use)
	}

	return users
}
