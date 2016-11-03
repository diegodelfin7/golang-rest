package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Move struct {
	Title  string `json:"title"`
	Rating string `json:"rating"`
	Year   string `json:"year"`
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", Index).Methods("GET")
	r.HandleFunc("/hello/{name}", IndexHello).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func IndexHello(w http.ResponseWriter, r *http.Request) {
	log.Println("Response to /hello request")
	log.Println(r.UserAgent())

	vars := mux.Vars(r)
	name := vars["name"]

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello :", name)
}
