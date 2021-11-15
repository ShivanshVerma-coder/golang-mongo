package router

import (
	controller "github.com/ShivanshVerma-coder/golang-mongo/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", controller.Welcome).Methods("GET")
	router.HandleFunc("/api/movies", controller.GetAllMoviesController).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovieController).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkMovieAsWatchedController).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteOneMovieController).Methods("DELETE")
	router.HandleFunc("/api/movies", controller.DeleteAllMovieController).Methods("DELETE")

	return router
}
