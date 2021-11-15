package controller

import (
	"encoding/json"
	"net/http"

	database "github.com/ShivanshVerma-coder/golang-mongo/database"
	model "github.com/ShivanshVerma-coder/golang-mongo/models"
	"github.com/gorilla/mux"
)

//Actual Controllers - thus the file name

func GetAllMoviesController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := database.GetAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func MarkMovieAsWatchedController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	movieId := params["id"]
	database.UpdateOneMovie(movieId)
	json.NewEncoder(w).Encode("Movie marked as watched")
}

func CreateMovieController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	database.InsertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func DeleteOneMovieController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	movieId := params["id"]
	deleteCount := database.DeleteOneMovie(movieId)
	json.NewEncoder(w).Encode(deleteCount)
}

func DeleteAllMovieController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	result := database.DeleteAllMovies("")
	json.NewEncoder(w).Encode(result)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Server"))
}
