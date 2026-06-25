package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"` 
	Phone string `json:"phone"`
}

var contacts []Contact
var nextID = 1

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}

func getContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

func addContact(w http.ResponseWriter, r *http.Request) {

	var c Contact

	err := json.NewDecoder(r.Body).Decode(&c)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c.ID = nextID
	nextID++

	contacts = append(contacts, c)

	json.NewEncoder(w).Encode(c)
}

func updateContact(w http.ResponseWriter, r *http.Request) {

	var updated Contact

	json.NewDecoder(r.Body).Decode(&updated)

	for i := range contacts {

		if contacts[i].ID == updated.ID {
			contacts[i] = updated
			break
		}
	}

	json.NewEncoder(w).Encode(updated)
}

func deleteContact(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")

	id, _ := strconv.Atoi(idStr)

	for i, c := range contacts {

		if c.ID == id {

			contacts = append(
				contacts[:i],
				contacts[i+1:]...,
			)

			break
		}
	}

	w.Write([]byte("Deleted"))
}

func main() {

	http.HandleFunc("/", home)

	http.HandleFunc("/contacts", getContacts)

	http.HandleFunc("/add", addContact)

	http.HandleFunc("/update", updateContact)

	http.HandleFunc("/delete", deleteContact)

	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	http.ListenAndServe(":8080", nil)
}