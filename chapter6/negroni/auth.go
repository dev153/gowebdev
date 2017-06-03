package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
)

func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token := r.Header.Get("X-AppToken")
	if token == "qwerty" {
		log.Println("Authorized to the system")
		context.Set(r, "user", "John")
		next(w, r)
	} else {
		http.Error(w, "Not Authorized", http.StatusUnauthorized)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	user, ok := context.GetOk(r, "user")
	if ok == true {
		fmt.Fprintf(w, "Welcome %s!", user)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(Authorize))
	n.UseHandler(mux)
	n.Run(":9099")
}
