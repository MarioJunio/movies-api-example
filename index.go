package main

import (
	"fmt"
	"log"
	"movies-api/resource"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Use(commonMiddleware)

	router.HandleFunc("/movies/", resource.GetMovies).Methods("GET")
	router.HandleFunc("/movies/", resource.CreateMovie).Methods("POST")
	router.HandleFunc("/movies/{movieId}", resource.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies/{movieId}", resource.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/movies/", resource.DeleteMovies).Methods("DELETE")

	fmt.Println("Server starting at port: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
