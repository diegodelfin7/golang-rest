package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	Title  string `json:"title"`
	Rating string `json:"rating"`
	Year   string `json:"year"`
}

// tabla hash map [keyType]ValueType
// tabla hash map [string] punteroMovie
var movies = map[string]*Movie{
	"tt00444": &Movie{Title: "Star Wars : A New Hope", Rating: "9.9", Year: "2000"},
	"tt00333": &Movie{Title: "Indiana Jones : Raider of the Lost Ark", Rating: "8.5", Year: "1999"},
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/movies", handleMovies).Methods("GET")
	r.HandleFunc("/movies/{movieId}", handleSearchMovie).Methods("GET")

	http.ListenAndServe(":8080", r)

}

func handleSearchMovie(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	keyID := vars["movieId"]

	log.Println("Request for :", keyID)

	movie, ok := movies[keyID]
	fmt.Println("Movie", movie)
	fmt.Println("ok", ok)

	if ok {
		outgointJSON, err := json.Marshal(movie)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, string(outgointJSON))
	} else {
		var ErrorJson = `{"error":"NotFound Error,"code":"404"}`
		fmt.Fprint(w, ErrorJson)
		return
	}
}

func handleMovies(w http.ResponseWriter, r *http.Request) {

	r.Header.Set("Content-Type", "application/json")

	outgoinJson, err := json.Marshal(movies)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(outgoinJson))
}
