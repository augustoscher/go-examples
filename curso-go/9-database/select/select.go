package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id   int
	name string
}

func main() {
	db, error := sql.Open("mysql", "root:root@tcp(172.17.0.2:3306)/test")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Connection opened successfully!")
	defer db.Close()

	rows, _ := db.Query("select id, name from users where id > ?", 1)
	defer rows.Close()

	var slice []user

	for rows.Next() {
		var u user
		rows.Scan(&u.id, &u.name)
		fmt.Println(u)
		slice = append(slice, u)
	}

	fmt.Println(slice)
	fmt.Println(slice[0])

}
