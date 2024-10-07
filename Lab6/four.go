// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var counter int
// 	var mutex sync.Mutex
// 	var wg sync.WaitGroup
// 	numGoroutines := 5
// 	incrementsPerGoroutine := 1000

// 	useMutex := true

// 	increment := func() {
// 		defer wg.Done()
// 		for i := 0; i < incrementsPerGoroutine; i++ {
// 			if useMutex {
// 				mutex.Lock()
// 			}

// 			counter++

// 			if useMutex {
// 				mutex.Unlock()
// 			}
// 		}
// 	}

// 	wg.Add(numGoroutines)
// 	for i := 0; i < numGoroutines; i++ {
// 		go increment()
// 	}

// 	wg.Wait()

// 	fmt.Printf("Финальное значение счётчика: %d\n", counter)
// }
