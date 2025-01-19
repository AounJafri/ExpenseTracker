package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Aounjafri/ExpenseTracker/auth"
	"github.com/Aounjafri/ExpenseTracker/models"
	"golang.org/x/crypto/bcrypt"
)

// var currentUserId int = 0

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	er := models.CreateUser(user.Name, string(hash), user.Email)

	if er != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(er.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User successfully created")
}

var CurrentUserToken string

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	// fmt.Println(user)

	uId, hash := models.GetHashedPassAndId(user.Name)

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(user.Password))

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Password is incorrect")
		return
	}

	// currentUserId = uId

	token, err := auth.GenerateToken(uId)
	CurrentUserToken = token

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Can't generate Token")
		return
	}

	fmt.Println("currently logged in user's id is :", uId)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("Authorization")

	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Missing Token")
		return
	}

	token = token[len("Bearer "):]

	if token == CurrentUserToken {

		tr, _ := auth.VerifyToken(token)

		if tr {
			CurrentUserToken = ""
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Logout Successfull")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("You are not the logged in user")
	}

}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users := models.GetUsers()

	json.NewEncoder(w).Encode(users)

}
