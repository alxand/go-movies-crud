package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isdn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lasttname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func createMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
		}
	}
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "2", Isbn: "29932", Title: "Hitman", Director: &Director{Firstname: "William", Lasttname: "Amoah"}})
	movies = append(movies, Movie{ID: "1", Isbn: "23232", Title: "Colman", Director: &Director{Firstname: "Alexan", Lasttname: "Adade"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Printf("Server is starting on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
