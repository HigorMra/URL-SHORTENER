package main

import (
	"fmt"
	"log"
	"net/http"
	"url-shortener/handler" // Importa o nosso pacote local
)

func main() {
	const port = ":8080"
	
	// Inicializa o armazenamento de teste no nosso pacote handler
	handler.InitStore()

	// 1. Rota para criar novas URLs
	http.HandleFunc("/shorten", handler.ShortenURLHandler)

	// 2. Rota para redirecionar URLs (a rota principal '/')
	// Se o caminho não for /shorten, ele tenta redirecionar
	http.HandleFunc("/", handler.RedirectURLHandler)

	fmt.Printf("Servidor iniciado na porta %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
