package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type entry struct {
	ID           int    `json:"id,omitempty"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`
	EmailAddress string `json:"email_address,omitempty"`
}

type allEntries []entry

var entries = allEntries{
	{
		ID:           1,
		FirstName:    "Kan",
		LastName:     "Ouivirach",
		EmailAddress: "test@example.com",
	},
}

func GetEntries(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(entries)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	var router *mux.Router
	router = mux.NewRouter().StrictSlash(true)
	router.Use(commonMiddleware)

	router.HandleFunc("/entries", GetEntries).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
