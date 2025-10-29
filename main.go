package main

import (
	"fmt"
	"log"
	"net/http"
	"url-shortener/handler" 
)

func main() {
	const port = ":8080"
	
	handler.InitStore()

	http.HandleFunc("/shorten", handler.ShortenURLHandler)

	http.HandleFunc("/", handler.RedirectURLHandler)

	fmt.Printf("Servidor iniciado na porta %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

