package main

import (
	"criminal_report/config"
	"criminal_report/handlers"
	"criminal_report/repository"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.InitializeDB()
	defer config.CloseDB()

	repo := repository.NewCriminalRepository(config.DB)
	handler := handlers.NewCriminalHandler(repo)

	router := mux.NewRouter()

	router.HandleFunc("/criminals", handler.GetCriminals).Methods("GET")
	router.HandleFunc("/criminals/{id}", handler.GetCriminal).Methods("GET")
	router.HandleFunc("/criminals", handler.CreateCriminal).Methods("POST")
	router.HandleFunc("/criminals/{id}", handler.UpdateCriminal).Methods("PUT")
	router.HandleFunc("/criminals/{id}", handler.DeleteCriminal).Methods("DELETE")

	fmt.Println("Server is listening on http://127.0.0.1:8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
