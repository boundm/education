package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Ошибка при апгрейде до WebSocket:", err)
		return
	}
	defer ws.Close()

	// Обработка сообщений от клиента
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("Ошибка чтения сообщения:", err)
			break
		}
		log.Printf("Получено сообщение: %s", msg)
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)

	log.Println("Запуск сервера WebSocket на порту 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
