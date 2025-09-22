package main

import (
	"github.com/gorilla/mux"
	"log"
	"main/pkg/db"
	"main/pkg/handlers"
	"net/http"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/news", h.GetAllNews).Methods(http.MethodGet)
	router.HandleFunc("/news/{id}", h.GetNew).Methods(http.MethodGet)
	router.HandleFunc("/contacts", h.GetContacts).Methods(http.MethodGet)
	router.HandleFunc("/blocks/{code}", h.GetBlocks).Methods(http.MethodGet)

	fs := http.FileServer(http.Dir("./public/"))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}
