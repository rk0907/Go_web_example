package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Handler for requests with HTTPS scheme
	r.HandleFunc("/secure", SecureHandler).Schemes("https")

	// Handler for requests with HTTP scheme
	r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

	// Start the HTTP server
	http.ListenAndServe(":80", r)
	http.ListenAndServe(":443", r)
}

// SecureHandler handles requests with HTTPS scheme
func SecureHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a secure handler (HTTPS)")
}

// InsecureHandler handles requests with HTTP scheme
func InsecureHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is an insecure handler (HTTP)")
}
