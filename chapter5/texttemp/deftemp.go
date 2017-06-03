package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	t, err := template.New("deftemp").Parse(`{{define "T"}}{{$cruel:="cruel"}}Hello, {{$cruel}} {{.}}!{{end}}`)
	err = t.ExecuteTemplate(os.Stdout, "T", "World")
	if err != nil {
		log.Fatal("Execute: ", err)
		return
	}
}
