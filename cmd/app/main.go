package main

import (
	"log"
	"net/http"

	"cicd/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler.Health)

	log.Println("server start :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
