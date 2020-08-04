package main

import (
	"encoding/json"
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

var noteStore = make(map[string]Note)

var id int = 0

//POST
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)

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

//GET
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	var notes []Note
	for _, value := range noteStore {
		notes = append(notes, value)
	}

	j, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes", GetNoteHandler).Methods("GET")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Printf("Listening")
	server.ListenAndServe()
}
