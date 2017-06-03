package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedOn   time.Time `json:"createdon"`
}

//Store for the Notes collecion
var noteStore = make(map[string]Note)

//Variable to generate key for the collection
var id int = 0

func PostNoteHandler(w http.ResponseWriter, req *http.Request) {
	var note Note
	err := json.NewDecoder(req.Body).Decode(&note)
	if err != nil {
		panic(err)
	}
	note.CreatedOn = time.Now()
	id++
	k := strconv.Itoa(id)
	noteStore[k] = note
	j, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func GetNoteHandler(w http.ResponseWriter, req *http.Request) {
	var notes []Note
	for _, note := range noteStore {
		notes = append(notes, note)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func PutNoteHandler(w http.ResponseWriter, req *http.Request) {
	var err error
	requestVars := mux.Vars(req)
	k := requestVars["id"]
	var noteToUpdate Note
	err = json.NewDecoder(req.Body).Decode(&noteToUpdate)
	if err != nil {
		panic(err)
	}
	if note, ok := noteStore[k]; ok {
		noteToUpdate.CreatedOn = note.CreatedOn
		// delete the existing item and add the updated item
		delete(noteStore, k)
		noteStore[k] = noteToUpdate
	} else {
		log.Printf("Could not find key of Note %s to put", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

func DeleteNoteHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	k := vars["id"]
	if _, ok := noteStore[k]; ok {
		delete(noteStore, k)
	} else {
		log.Printf("Could not find key of Note %s to delete", k)
	}
	w.WriteHeader(http.StatusNoContent)
}

func GetNoteByIdHandler(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	noteID := vars["id"]

	if noteToSend, ok := noteStore[noteID]; !ok {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("<html><head></head><body>Could not find Note with Id: %s</body></html>", noteID)))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		jsonReply, err := json.Marshal(noteToSend)
		if err != nil {
			panic(err)
		} else {
			w.Write(jsonReply)
		}
	}
}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes/{id}", GetNoteByIdHandler).Methods("GET")
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", DeleteNoteHandler).Methods("DELETE")

	server := &http.Server{
		Addr:    ":9099",
		Handler: r,
	}
	log.Println("Listening...")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
