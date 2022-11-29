package routes

import (
	"github.com/gorilla/mux"
	"github.com/rajatk31/mongo/controllers"
)

func Routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controllers.FindAll).Methods("GET")
	router.HandleFunc("/api/movie", controllers.CreateMovies).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controllers.MarkWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controllers.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteallmovie", controllers.DeletteAll).Methods("DELETE")

	return router
}
