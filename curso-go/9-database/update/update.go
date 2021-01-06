package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, error := sql.Open("mysql", "root:root@tcp(172.17.0.2:3306)/test")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Connection opened successfully!")
	defer db.Close()

	stmt, _ := db.Prepare("update users set name = ? where id = ?")
	stmt.Exec("Update", 1)
	stmt.Exec("Schersnegger", 2)
}
