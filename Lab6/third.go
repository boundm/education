// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func main() {
// 	numberChannel := make(chan int)
// 	resultChannel := make(chan string)
// 	done := make(chan bool)

// 	go func() {
// 		for {
// 			num := rand.Intn(100)
// 			numberChannel <- num
// 			time.Sleep(time.Millisecond * 500)
// 		}
// 	}()

// 	go func() {
// 		for {
// 			num := <-numberChannel
// 			if num%2 == 0 {
// 				resultChannel <- fmt.Sprintf("Число %d - чётное", num)
// 			} else {
// 				resultChannel <- fmt.Sprintf("Число %d - нечётное", num)
// 			}
// 		}
// 	}()

// 	go func() {
// 		count := 0
// 		for {
// 			select {
// 			case result := <-resultChannel:
// 				fmt.Println(result)
// 				count++
// 				if count >= 10 {
// 					done <- true
// 					return
// 				}
// 			}
// 		}
// 	}()
// 	<-done
// 	fmt.Println("Завершение программы.")
// }
