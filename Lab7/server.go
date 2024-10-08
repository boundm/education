// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"net"
// 	"os"
// 	"os/signal"
// 	"sync"
// 	"syscall"
// )

// var wg sync.WaitGroup

// func handleConnection(conn net.Conn) {
// 	defer wg.Done()
// 	defer conn.Close()

// 	reader := bufio.NewReader(conn)
// 	message, err := reader.ReadString('\n')
// 	if err != nil {
// 		log.Println("Ошибка при чтении от клиента:", err)
// 		return
// 	}

// 	fmt.Printf("Получено сообщение от клиента: %s", message)

// 	_, err = conn.Write([]byte("Сообщение получено!\n"))
// 	if err != nil {
// 		log.Println("Ошибка при отправке подтверждения клиенту:", err)
// 		return
// 	}
// }

// func main() {
// 	listener, err := net.Listen("tcp", ":8080")
// 	if err != nil {
// 		log.Fatal("Ошибка запуска TCP-сервера:", err)
// 	}
// 	defer listener.Close()
// 	log.Println("TCP-сервер запущен на порту 8080")

// 	stopChan := make(chan os.Signal, 1)
// 	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

// 	go func() {
// 		<-stopChan
// 		log.Println("Сервер завершает работу...")
// 		listener.Close()
// 	}()

// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			select {
// 			case <-stopChan:
// 				log.Println("Сервер остановлен, завершение всех соединений...")
// 				wg.Wait()
// 				log.Println("Все соединения завершены")
// 				os.Exit(0)
// 			default:
// 				log.Println("Ошибка при принятии соединения:", err)
// 			}
// 			continue
// 		}

// 		wg.Add(1)
// 		go handleConnection(conn)
// 	}
// }
