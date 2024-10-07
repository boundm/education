package main

import (
	"fmt"
	"sync"
)

func generateFibonacci(n int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch) // закрытие канала после завершения функции

	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a //отправляем в канал
		a, b = b, a+b
	}
}

func printFibonacci(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch {
		fmt.Println(num)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int) //создание канала

	wg.Add(2)
	go generateFibonacci(10, ch, &wg)
	go printFibonacci(ch, &wg)

	wg.Wait()
	fmt.Println("все функции завершились")
}
