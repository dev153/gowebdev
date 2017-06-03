package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tmpl, err := template.New("test-security").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	err = tmpl.ExecuteTemplate(os.Stdout, "T", "<script>alert('XSS Injection')</script>")
	if err != nil {
		log.Fatal("Execute: ", err)
	}
}
