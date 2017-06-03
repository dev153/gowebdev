package main

import (
	"fmt"
	"log"
	"net/http"
)

func messageHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to Go Web Development")
}

func aboutMessageHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "About Go Web Development")
}

func main() {
	mux := http.NewServeMux()
	// convert the messageHandler function to a HandlerFunc type.
	mh := http.HandlerFunc(messageHandler)
	mux.Handle("/welcome", mh)
	aboutMsgHndl := http.HandlerFunc(aboutMessageHandler)
	mux.Handle("/about", aboutMsgHndl)
	log.Println("Listening...")
	http.ListenAndServe(":9099", mux)
}
