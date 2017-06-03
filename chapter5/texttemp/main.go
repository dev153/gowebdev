package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type Note struct {
	Title       string
	Description string
}

const tmpl = `Note - Title: {{.Title}}, Description: {{.Description}}`

func main() {
	note := Note{
		Title:       "text/templates",
		Description: "Text-Templates generate textual output",
	}

	t := template.New("note")
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}
	// Appies a parsed template to the data of Note object
	if err := t.Execute(os.Stdout, note); err != nil {
		log.Fatal("Execute: ", err)
		return
	}
	fmt.Println()
	fmt.Println((*t).Name())
}
