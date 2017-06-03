package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/pat"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
)

type Configuration struct {
	FacebookKey    string
	FacebookSecret string
}

var config Configuration

var indexTemplate = `
<p><a href="/auth/facebook">Log in with Facebook</a></p>
`

var userTemplate = `
<p>Name: {{.Name}}</p>
<p>Email: {{.Email}}
`

func init() {
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
}

func callbackAuthHandler(res http.ResponseWriter, req *http.Request) {
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		fmt.Fprintln(res, err)
		return
	}
	tmpTemplate, _ := template.New("userinfo").Parse(userTemplate)
	tmpTemplate.Execute(res, user)
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	tmpTemplate, _ := template.New("index").Parse(indexTemplate)
	tmpTemplate.Execute(res, nil)
}

func main() {
	fmt.Println(config)
	goth.UseProviders(
		facebook.New(config.FacebookKey, config.FacebookSecret, "http://localhost:9099/auth/facebook/callback"),
	)

	router := pat.New()
	router.Get("/auth/{provider}/callback", callbackAuthHandler)
	router.Get("/auth/{provider}", gothic.BeginAuthHandler)
	router.Get("/", indexHandler)

	server := &http.Server{
		Addr:    ":9099",
		Handler: router,
	}
	server.ListenAndServe()
}
