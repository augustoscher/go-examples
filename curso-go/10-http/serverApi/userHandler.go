package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// User struct
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// UserHandler analyze the request and delegate to function
func UserHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		getUserByID(w, r, id)
	case r.Method == "GET":
		getAllUsers(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Sorry :(")
	}
}

func getUserByID(w http.ResponseWriter, r *http.Request, id int) {
	db, error := sql.Open("mysql", "root:root@tcp(172.17.0.2:3306)/test")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Connection opened successfully!")
	defer db.Close()

	var u User
	db.QueryRow("select id, name from users where id = ?", id).Scan(&u.ID, &u.Name)
	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	db, error := sql.Open("mysql", "root:root@tcp(172.17.0.2:3306)/test")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Connection opened successfully!")
	defer db.Close()

	rows, _ := db.Query("select id, name from users")
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Name)
		users = append(users, user)
	}

	json, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}
