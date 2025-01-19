package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Aounjafri/ExpenseTracker/controllers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/logout", controllers.Logout).Methods("GET")

	r.HandleFunc("/expense", controllers.CreateExpense).Methods("POST")
	r.HandleFunc("/expenses", controllers.GetExpenses).Methods("GET")
	r.HandleFunc("/expense/{expenseid}", controllers.GetExpense).Methods("GET")
	r.HandleFunc("/expense/{expenseid}", controllers.UpdateExpense).Methods("PUT")
	r.HandleFunc("/expense/{expenseid}", controllers.DeleteExpense).Methods("DELETE")

	http.Handle("/", r)
	fmt.Println("Server starting at port 8080.....")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
