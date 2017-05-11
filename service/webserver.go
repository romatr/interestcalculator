package service

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartWebServer(port string) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/calculate/{name}", handleCalculate).Methods("GET")

	log.Println("Starting HTTP service at " + port)
	log.Fatal(http.ListenAndServe(":" + port, router))
}

