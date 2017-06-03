package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Note is the data type on which the CRUD operations
//of the web application will applied on.
type Note struct {
	Title       string
	Description string
	CreatedOn   time.Time
}

//EditNote is the structure for updating an existing Note.
type EditNote struct {
	Note
	ID string
}

//A collection to store the "Note" structures' data.
var noteStore = make(map[string]Note)

//A variable that will assign a unique id to each new "Note" data.
var noteID = 0

const serverAddress string = ":9099"

var templates map[string]*template.Template

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	templates["index"] = template.Must(template.ParseFiles("templates/index.html", "templates/base.html"))
	templates["add"] = template.Must(template.ParseFiles("templates/add.html", "templates/base.html"))
	templates["edit"] = template.Must(template.ParseFiles("templates/edit.html", "templates/base.html"))
}

//Render templates for the given name, template definition and data object.
func renderTemplate(w http.ResponseWriter, name string, template string, viewModel interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "The template does not exist", http.StatusInternalServerError)
	}
	err := tmpl.ExecuteTemplate(w, template, viewModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//=============================================================================

func getNotes(w http.ResponseWriter, req *http.Request) {
	fmt.Println("getNotes()")
	renderTemplate(w, "index", "base", noteStore)
}

func addNote(w http.ResponseWriter, req *http.Request) {
	fmt.Println("addNote()")
	renderTemplate(w, "add", "base", nil)
}

func saveNote(w http.ResponseWriter, req *http.Request) {
	fmt.Println("saveNote()")
	// Parse the form in order to obtain the arguments sent.
	req.ParseForm()
	title := req.PostFormValue("title")
	description := req.PostFormValue("description")
	createdOn := time.Now()
	noteToAdd := Note{title, description, createdOn}
	fmt.Println(noteToAdd)
	/* increment the unique id of the note, convert it to string
	and store the note in the map.
	*/
	noteID++
	key := strconv.Itoa(noteID)
	noteStore[key] = noteToAdd
	http.Redirect(w, req, "/", http.StatusFound)
}

func editNote(w http.ResponseWriter, req *http.Request) {
	fmt.Println("editNote()")
	var viewModel EditNote
	vars := mux.Vars(req)
	index := vars["id"]
	fmt.Println("Note index to update", index)
	if note, ok := noteStore[index]; ok {
		viewModel = EditNote{note, index}
	} else {
		http.Error(w, "Could not find the resource to edit.", http.StatusBadRequest)
	}
	renderTemplate(w, "edit", "base", viewModel)
}

func updateNote(w http.ResponseWriter, req *http.Request) {
	fmt.Println("updateNote")
	vars := mux.Vars(req)
	index := vars["id"]
	var noteToUpdate Note
	if note, ok := noteStore[index]; ok {
		req.ParseForm()
		noteToUpdate.Title = req.PostFormValue("title")
		noteToUpdate.Description = req.PostFormValue("description")
		noteToUpdate.CreatedOn = note.CreatedOn
		delete(noteStore, index)
		noteStore[index] = noteToUpdate
	} else {
		http.Error(w, "Could not find the resource to update.", http.StatusBadRequest)
	}
	http.Redirect(w, req, "/", http.StatusFound)
}

func deleteNote(w http.ResponseWriter, req *http.Request) {
	fmt.Println("deleteNote")
	vars := mux.Vars(req)
	index := vars["id"]
	fmt.Println("Note index to delete", index)
	if _, ok := noteStore[index]; ok {
		delete(noteStore, index)
	} else {
		http.Error(w, "Could not find the resource to delete.", http.StatusBadRequest)
	}
	http.Redirect(w, req, "/", http.StatusFound)
}

//=============================================================================

func main() {
	// Creating the route multiplexer to be used by the web application.
	router := mux.NewRouter().StrictSlash(false)
	fs := http.FileServer(http.Dir("public"))
	router.Handle("/public/", fs)
	router.HandleFunc("/", getNotes)
	router.HandleFunc("/notes/add", addNote)
	router.HandleFunc("/notes/save", saveNote)
	router.HandleFunc("/notes/edit/{id}", editNote)
	router.HandleFunc("/notes/update/{id}", updateNote)
	router.HandleFunc("/notes/delete/{id}", deleteNote)
	server := &http.Server{
		Addr:    serverAddress,
		Handler: router,
	} // do not forget on a multiline initialization to add the last comma.

	log.Printf("Listening...\n")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}
