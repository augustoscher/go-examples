package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	log.Println("Running...")
	log.Fatal(http.ListenAndServe(":3030", nil))
}
