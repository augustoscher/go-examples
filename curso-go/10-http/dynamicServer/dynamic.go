package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/rightTime", rightTime)
	log.Println("Running...")
	log.Fatal(http.ListenAndServe(":3030", nil))
}

func rightTime(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05") //string de formatação
	fmt.Fprintf(w, "<h1>Right Time: %s</h1>", s)
}
