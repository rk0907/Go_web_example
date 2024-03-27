package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	//createBook handler
	r.HandleFunc("/books/{title}", CreatBook).Methods("POST")

	//ReadBook Handler
	r.HandleFunc("/books/{title}", ReadBook).Methods("GET")

	//UpdateBook handler
	r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")

	//DeleteBook Handler
	r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

	http.ListenAndServe(":80", r)

}

func CreatBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	fmt.Fprintf(w, "CreateBook: You've created a book with title: %s\n", title)
}

func ReadBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	fmt.Fprintf(w, "ReadBook: You've read a book with title: %s\n", title)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	fmt.Fprintf(w, "UpdateBook: You've updated a book with title: %s\n", title)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	fmt.Fprintf(w, "Deleteook: You've Deleted a book with title: %s\n", title)
}
