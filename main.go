package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	// Code to shorten the URL
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
