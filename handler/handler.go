package handler

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

var urlStore = make(map[string]string)

func InitStore() {
	
	urlStore["go-docs"] = "https://go.dev/doc/"
	urlStore["github"] = "https://github.com/google/go-cloud"
	fmt.Println("URLs de teste carregadas.")
}

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

func ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	newShortKey := "go-home"
	newLongURL := "https://go.dev/"

	urlStore[newShortKey] = newLongURL

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Nova URL criada!\nCurta: /%s\nLonga: %s\n", newShortKey, newLongURL)))
}
