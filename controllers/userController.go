package controllers

import (
	"crud/database"
	"crud/entities"
	"crud/repositories/userRepositories"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateUser insert a new user into the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Failed to read body"))
		return
	}

	var user entities.User

	if err = json.Unmarshal(reqBody, &user); err != nil {
		w.Write([]byte("Failed to parse user to struct"))
		return
	}

	result, err := database.ExecuteStatement("insert into users (name, email) values (?, ?)", user.Name, user.Email)
	if err != nil {
		w.Write([]byte("Failed to execute query: " + err.Error()))
	}
	insertedId, err := result.LastInsertId()
	if err != nil {
		w.Write([]byte("Failed to get rows affected"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User Id: %v inserted successfully", insertedId)))
}

// FindUsers returns a list of users saved in the database
func FindUsers(w http.ResponseWriter, r *http.Request) {
	var users, err = userRepositories.FindMany()
	if err != nil {
		w.Write([]byte("Failed to find users"))
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Failed to encode users to json"))
		return
	}
}

// FindUser finds a specific user in the database
func FindUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Failed to parse param"))
		return
	}

	user, err := userRepositories.FindById(ID)
	if err != nil {
		w.Write([]byte("Failed to find users"))
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Failed to encode user to JSON"))
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Failed to parse id param to int"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Failed to read body"))
		return
	}

	var user entities.User
	if err := json.Unmarshal(body, &user); err != nil {
		w.Write([]byte("Failed to parse user to struct"))
		return
	}
	
	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Failed to connect to database"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("update users set name = ?, email = ? where id = ?")
	if err != nil {
		w.Write([]byte("Failed to prepare statement"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Email, ID); err != nil {
		w.Write([]byte("Failed to update user"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Failed to conver param to int"))
		return
	}

	if _, err := database.ExecuteStatement("delete from users where id=?", ID); err != nil {
		w.Write([]byte("Failed to delete a user"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}