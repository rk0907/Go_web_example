package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Define a route that matches the "/books/{title}" pattern
	// and restricts it to requests with the hostname "www.mybookstore.com"
	r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

	// Start the HTTP server
	http.ListenAndServe(":80", r)
}

// BookHandler handles requests to the "/books/{title}" route
func BookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	fmt.Fprintf(w, "Handler for www.mybookstore.com: You've requested the book: %s\n", title)
}
