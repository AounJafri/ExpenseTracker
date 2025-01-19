package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Aounjafri/ExpenseTracker/auth"
	"github.com/Aounjafri/ExpenseTracker/models"
	"github.com/gorilla/mux"
)

// EXPENSES CONTROLLERS

func CreateExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("Authorization")

	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Missing Token")
		return
	}

	token = token[len("Bearer "):]

	if token == CurrentUserToken {

		tr, uId := auth.VerifyToken(token)

		if tr {
			var newexp models.Expense
			json.NewDecoder(r.Body).Decode(&newexp)

			id, err := models.CreateExpense(int(uId), newexp.Amount, newexp.Description, newexp.Category, newexp.Date)

			if err != nil {
				fmt.Println(err.Error())
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode("Error Creating expense")
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(models.GetExpense(id, int(uId)))

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Token invalid or expired, try logging in again")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("User token is either expired or user is not logged in")
	}
}

func GetExpenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("Authorization")

	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Missing Token")
		return
	}

	token = token[len("Bearer "):]

	if token == CurrentUserToken {
		tr, uId := auth.VerifyToken(token)

		if tr && uId != 0 {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(models.GetExpenses(int(uId)))
			return
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Token invalid or expired, try logging in again")
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("User token is either expired or user is not logged in")
	}
}

func GetExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("Authorization")

	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Missing Token")
		return
	}

	token = token[len("Bearer "):]

	if token == CurrentUserToken {

		tr, uId := auth.VerifyToken(token)

		if tr {

			ID, err := strconv.ParseInt(mux.Vars(r)["expenseid"], 0, 0)

			if err != nil {
				fmt.Println("Error parsing ID")
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(models.GetExpense(int(ID), int(uId)))

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Token invalid or expired, try logging in again")
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("User token is either expired or user is not logged in")
	}
}

func UpdateExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("Authorization")

	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Missing Token")
		return
	}

	token = token[len("Bearer "):]

	if token == CurrentUserToken {

		tr, uId := auth.VerifyToken(token)

		if tr {

			ID, err := strconv.ParseInt(mux.Vars(r)["expenseid"], 0, 0)

			if err != nil {
				fmt.Println("Error parsing ID")
			}

			var updExp models.Expense

			json.NewDecoder(r.Body).Decode(&updExp)

			retId, er := models.UpdateExpense(int(ID), int(uId), updExp)

			if er != nil {
				fmt.Println(er.Error())
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode("Error Updating expense")
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(models.GetExpense(retId, int(uId)))

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Token invalid or expired, try logging in again")
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("User token is either expired or user is not logged in")
	}
}

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("Authorization")

	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Missing Token")
		return
	}

	token = token[len("Bearer "):]

	if token == CurrentUserToken {

		tr, uId := auth.VerifyToken(token)

		if tr {

			ID, err := strconv.ParseInt(mux.Vars(r)["expenseid"], 0, 0)

			if err != nil {
				fmt.Println("Error parsing ID")
			}

			er := models.DeleteExpense(int(ID), int(uId))

			if er != nil {
				fmt.Println(er.Error())
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode("Error Deleting expense")
				return
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deletion Successful")

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Token invalid or expired, try logging in again")
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("User token is either expired or user is not logged in")
	}
}
