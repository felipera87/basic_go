package server

import (
	"crud/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUser create an user on the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Fail to read request body"))
		return
	}

	var user user

	if err = json.Unmarshal(requestBody, &user); err != nil {
		w.Write([]byte("Fail to convert JSON to user"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Couldn't connect to database"))
		return
	}
	defer db.Close()

	// this is used to avoid sql injection
	statement, err := db.Prepare("insert into users (name, email) values (?, ?)")
	if err != nil {
		w.Write([]byte("Couldn't create sql statement"))
		return
	}
	defer statement.Close()

	insertResult, err := statement.Exec(user.Name, user.Email)
	if err != nil {
		w.Write([]byte("Couldn't create user on database"))
		return
	}

	insertedID, err := insertResult.LastInsertId()
	if err != nil {
		w.Write([]byte("Couldn't get the created user id"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User successfully created, id=%d", insertedID)))
}

// SearchUsers get all users saved on database
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Couldn't connect to database"))
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from users")
	if err != nil {
		w.Write([]byte("Couldn't fetch users from database."))
		return
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var user user

		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Error when scanning user"))
			return
		}

		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Error to encode users to json"))
		return
	}
}

// SearchUser get one specified user
func SearchUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Error when converting parameter to integer."))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Couldn't connect to database"))
		return
	}
	defer db.Close()

	row, err := db.Query("select * from users where id = ?", ID)
	if err != nil {
		w.Write([]byte("Couldn't fetch user from database."))
		return
	}
	defer row.Close()

	var user user
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Error when scanning user"))
			return
		}
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Error to encode user to json"))
		return
	}
}

// UpdateUser updates one specified user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Error when converting parameter to integer."))
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Fail to read request body"))
		return
	}

	var user user

	if err = json.Unmarshal(requestBody, &user); err != nil {
		w.Write([]byte("Fail to convert JSON to user"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Couldn't connect to database"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("update users set name = ?, email = ? where id = ?")
	if err != nil {
		w.Write([]byte("Couldn't create sql statement"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Email, ID); err != nil {
		w.Write([]byte("Couldn't update user on database"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteUser deletes one specified user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Error when converting parameter to integer."))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Couldn't connect to database"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("delete from users where id = ?")
	if err != nil {
		w.Write([]byte("Couldn't create sql statement"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		w.Write([]byte("Couldn't delete user on database"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
