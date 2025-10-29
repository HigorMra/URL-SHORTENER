package handler

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// storage simula um banco de dados usando um mapa em memória
var urlStore = make(map[string]string)

// InitStore adiciona alguns dados de teste
func InitStore() {
	// Chave Curta: URL Longa
	urlStore["go-docs"] = "https://go.dev/doc/"
	urlStore["github"] = "https://github.com/google/go-cloud"
	fmt.Println("URLs de teste carregadas.")
}

// RedirectURLHandler lida com a requisição de redirecionamento (GET /<short-url>)
func RedirectURLHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := strings.TrimPrefix(r.URL.Path, "/")

	longURL, ok := urlStore[shortURL]

	if !ok {
		http.Error(w, "URL Curta não encontrada", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, longURL, http.StatusMovedPermanently)
	fmt.Printf("[%s] Redirecionamento para: %s\n", time.Now().Format("15:04:05"), longURL)
}

// ShortenURLHandler lida com a requisição de criação (POST /shorten)
func ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Simplificação: vamos apenas adicionar um novo par fixo
	newShortKey := "go-home"
	newLongURL := "https://go.dev/"

	urlStore[newShortKey] = newLongURL

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Nova URL criada!\nCurta: /%s\nLonga: %s\n", newShortKey, newLongURL)))
}
