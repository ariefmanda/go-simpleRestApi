package main

import (
	"log"
	"net/http"
	"./controllers"
	"github.com/gorilla/mux"
	"./models"
)

func main() {
	models.Connect()
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.AllMoviesEndPoint).Methods("GET")
	r.HandleFunc("/", controllers.CreateMovieEndPoint).Methods("POST")
	r.HandleFunc("/{id}", controllers.UpdateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/{id}", controllers.DeleteMovieEndPoint).Methods("DELETE")
	r.HandleFunc("/{id}", controllers.FindMovieEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}