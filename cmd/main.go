package main

import (
	"devops-toolbelt/api"
	"log"
	"net/http"
)

func main() {
	router := api.NewRouter()
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
