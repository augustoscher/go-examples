package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, error := sql.Open("mysql", "root:root@tcp(172.17.0.2:3306)/test")
	if error != nil {
		panic(error)
	}
	fmt.Println("Connection opened successfully!")
	defer db.Close()

	exec(db, "create database if not exists test")
	exec(db, "use test")
	exec(db, "drop table if exists users")
	exec(db, `create table users (
		id integer auto_increment primary key,
		name varchar(100) )`)
}

func exec(db *sql.DB, sql string) sql.Result {
	result, error := db.Exec(sql)
	if error != nil {
		panic(error)
	}
	return result
}
