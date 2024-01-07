package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type User struct {
	Username string `json: "username"`
	Email    string `json: "email"`
	Password string `json: "password"`
}

var users []User

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	users = append(users, newUser)

	// w.WriteHeader(http.StatusCreated)
	// fmt.Fprintf(w, "User registered successfully: %s", newUser.Username)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)

}

func GetRegisterUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Cotent-Type", "application/json")
	json.NewEncoder(w).Encode(users)

}

func GetUserByUsername(w http.ResponseWriter, r *http.Request) {

	username := strings.TrimPrefix(r.URL.Path, "/getUser/")

	for _, user := range users {
		if user.Username == username {
			w.Header().Set("Cotent-Type", "application/json")
			json.NewEncoder(w).Encode(users)
			return

		}

	}
	http.NotFound(w, r)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	username := strings.TrimPrefix(r.URL.Path, "/updateUser")
	var updatedUser User

	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	for i, user := range users {
		if user.Username == username {
			users[i] = updatedUser

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}
	http.NotFound(w, r)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	username := strings.TrimPrefix(r.URL.Path, "/deleteUser")

	for i, user := range users {
		if user.Username == username {
			users = append(users[:i], users[i+1:]...)

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

}

func main() {

	http.HandleFunc("/register", RegisterUser)
	http.HandleFunc("/getUsers", GetRegisterUser)
	http.HandleFunc("/getUser/", GetUserByUsername)
	http.HandleFunc("/updateUser", UpdateUser)
	http.HandleFunc("/deleteUser", DeleteUser)

	port := 8010

	fmt.Printf("server is runing on http://localhost:%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)

}
