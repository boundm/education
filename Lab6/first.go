// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// func factorial(n int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	result := 1
// 	for i := 1; i <= n; i++ {
// 		result *= i
// 		time.Sleep(100 * time.Millisecond)
// 	}
// 	fmt.Printf("Факториал %d = %d\n", n, result)
// }

// func generateRandomNumbers(count int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i < count; i++ {
// 		num := rand.Intn(100)
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Printf("Случайное число %d: %d\n", i+1, num)
// 	}
// }

// func sum(n int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	sum := 0
// 	for i := 1; i <= n; i++ {
// 		sum += i
// 		time.Sleep(100 * time.Millisecond)
// 	}
// 	fmt.Printf("Сумма числового ряда до %d = %d\n", n, sum)
// }

// func main() {
// 	var wg sync.WaitGroup

// 	wg.Add(3) //добавления кол-ва корутин в очередь
// 	go factorial(5, &wg)
// 	go generateRandomNumbers(3, &wg)
// 	go sum(5, &wg)

// 	wg.Wait() //ожидания завершения всех корутин

// 	fmt.Println("все функции завершились")
// }
