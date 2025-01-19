package models

import (
	"fmt"
)

type Expense struct {
	Id          int     `json:"id"`
	Amount      float64 `json:"amount"`
	UserId      int     `json:"userid"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Date        string  `json:"date"`
}

// EXPENSE FUNCTIONS

func CreateExpense(userId int, amount float64, descr, category, date string) (int, error) {

	result, err := db.Exec("INSERT INTO Expense(userid,amount,description,category,date_col) VALUES(?,?,?,?,?);",
		userId, amount, descr, category, date,
	)

	if err != nil {
		fmt.Println(err.Error())
	}

	id, _ := result.LastInsertId()

	return int(id), err
}

func GetExpense(id, userid int) Expense {
	var exp Expense

	db.QueryRow("Select * From Expense Where id = ? And userid=?", id, userid).Scan(&exp.Id, &exp.Amount, &exp.UserId, &exp.Description, &exp.Category, &exp.Date)
	// fmt.Println(exp)

	return exp
}

func GetExpenses(userId int) []Expense {
	result, err := db.Query("Select * From Expense Where userid = ?", userId)

	if err != nil {
		fmt.Println(err.Error())
	}

	var expenses []Expense

	for result.Next() {
		var exp Expense
		result.Scan(&exp.Id, &exp.Amount, &exp.UserId, &exp.Description, &exp.Category, &exp.Date)
		expenses = append(expenses, exp)
	}
	return expenses
}

func UpdateExpense(id, userid int, updatedExp Expense) (int, error) {
	var newExp Expense

	if updatedExp.Amount != 0.0 {
		newExp.Amount = updatedExp.Amount
	}
	if updatedExp.Description != "" {
		newExp.Description = updatedExp.Description
	}
	if updatedExp.Category != "" {
		newExp.Category = updatedExp.Category
	}
	if updatedExp.Date != "" {
		newExp.Date = updatedExp.Date
	}
	_, err := db.Exec("Update Expense SET amount = ?,category=?,description=?,date_col=? WHERE id =? AND userid=?", newExp.Amount, newExp.Category, newExp.Description, newExp.Date, id, userid)

	if err != nil {
		fmt.Println(err.Error())
	}

	return id, err
}

func DeleteExpense(id, uId int) error {
	_, err := db.Exec("Delete From Expense Where id = ? And userid=?", id, uId)

	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}
