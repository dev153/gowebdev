package main

import (
	"fmt"
	"log"
	"net/http"
)

func messageHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to Go Web Development")
}

func main() {
	http.HandleFunc("/welcome", messageHandler)
	log.Println("Listening...")
	http.ListenAndServe(":9099", nil)
}
