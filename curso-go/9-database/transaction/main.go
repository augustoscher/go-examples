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

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into users(id, name) values(?, ?)")
	stmt.Exec(200, "Oi")
	stmt.Exec(201, "Oi2")

	//simulating error
	_, err := stmt.Exec(1, "Oi3")

	if err != nil {
		tx.Rollback()
		log.Fatal(err) //faz o programa parar e nao executar o commit
	}

	tx.Commit()
}
