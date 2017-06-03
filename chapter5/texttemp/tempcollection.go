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

const collectionTmplStr = `Notes are:
{{range .}}
    Note - Title: {{.Title}}, Description: {{.Description}}
{{end}}
`

func main() {
	note1 := Note{
		Title:       "text/templates",
		Description: "Text-Templates generate textual output",
	}
	note2 := Note{
		Title:       "html/templates",
		Description: "Html-Templates generate html output",
	}

	var notes []Note
	notes = append(notes, note1, note2)

	t := template.New("note")
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}
	// Appies a parsed template to the data of Note object
	if err := t.Execute(os.Stdout, note1); err != nil {
		log.Fatal("Execute: ", err)
		return
	}
	fmt.Println()
	fmt.Println((*t).Name())

	collectionTmpl := template.New("Collection-Template")
	collectionTmpl, err = collectionTmpl.Parse(collectionTmplStr)
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := collectionTmpl.Execute(os.Stdout, notes); err != nil {
		log.Fatal(err)
		return
	}
}
