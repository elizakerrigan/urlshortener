package main

import (
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	// Code to shorten the URL
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing Url", http.StatusBadRequest)
		return
	}

	// get the URL from the form
	longURL := r.FormValue("url")
	if longURL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	id := uuid.New()
	fmt.Println(id.String())

    // Save the ID and URL in your storage system
    if err := saveURL(id, longURL); err != nil {
        http.Error(w, "Error saving URL", http.StatusInternalServerError)
        return
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	// Code to redirect the shortened URL
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/shorten", shortenURLHandler).Methods("POST")
	r.HandleFunc("/{id}", redirectHandler).Methods("GET")

	http.ListenAndServe(":8080", r)
}
