package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	// Создаём URL для подключения к серверу WebSocket
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
	fmt.Printf("Подключаемся к серверу: %s\n", u.String())

	// Подключение к WebSocket серверу
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Ошибка подключения к WebSocket:", err)
	}
	defer c.Close()

	// Канал для обработки прерывания (Ctrl+C) и корректного завершения
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Канал для отправки сообщений на сервер
	done := make(chan struct{})

	// Запускаем горутину для получения сообщений от сервера
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("Ошибка при чтении сообщения:", err)
				return
			}
			log.Printf("Получено сообщение: %s", message)
		}
	}()

	// Цикл для отправки сообщений
	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("Прерывание, закрытие соединения...")

			// Отправка закрывающего сообщения
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Ошибка при закрытии соединения:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		default:
			// Отправка сообщения каждые 5 секунд
			time.Sleep(5 * time.Second)
			message := "От клиента"
			err := c.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Println("Ошибка отправки сообщения:", err)
				return
			}
			log.Printf("Отправлено сообщение: %s", message)
		}
	}
}
