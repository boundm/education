// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"net"
// 	"os"
// )

// func main() {
// 	// Подключение к серверу
// 	conn, err := net.Dial("tcp", "localhost:8080")
// 	if err != nil {
// 		log.Fatal("Ошибка подключения к серверу:", err)
// 	}
// 	defer conn.Close()

// 	// Чтение сообщения от пользователя
// 	fmt.Print("Введите сообщение для сервера: ")
// 	reader := bufio.NewReader(os.Stdin)
// 	message, err := reader.ReadString('\n')
// 	if err != nil {
// 		log.Fatal("Ошибка чтения ввода:", err)
// 	}

// 	// Отправка сообщения серверу
// 	_, err = conn.Write([]byte(message))
// 	if err != nil {
// 		log.Fatal("Ошибка отправки сообщения серверу:", err)
// 	}

// 	// Ожидание ответа от сервера
// 	reply, err := bufio.NewReader(conn).ReadString('\n')
// 	if err != nil {
// 		log.Fatal("Ошибка при получении ответа от сервера:", err)
// 	}

// 	// Вывод ответа
// 	fmt.Printf("Ответ от сервера: %s", reply)
// }
