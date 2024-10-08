package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Middleware для логирования запросов
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

// Обработчик GET-запроса по пути /hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "Привет!")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Обработчик POST-запроса по пути /data
func dataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var data map[string]interface{}
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&data)
		if err != nil {
			http.Error(w, "неверный JSON", http.StatusBadRequest)
			return
		}
		fmt.Println("Received data:", data)
		fmt.Fprintln(w, "Data received successfully")
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	mux := http.NewServeMux()

	// Регистрация обработчиков
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/data", dataHandler)

	// Применение middleware для логирования
	loggedMux := loggingMiddleware(mux)

	log.Println("Запуск HTTP-сервера на порту 8080")
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}

//curl -X POST http://localhost:8080/data -H "Content-Type: application/json" -d "{\"name\": \"Dima\", \"age\": 99}"
//localhost:8080/hello

//go get github.com/gorilla/websocket
