// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// type CalcRequest struct {
// 	Operand1  float64
// 	Operand2  float64
// 	Operation string
// 	Result    chan float64
// }

// func calculator(requests chan CalcRequest, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for req := range requests {
// 		var result float64
// 		switch req.Operation {
// 		case "+":
// 			result = req.Operand1 + req.Operand2
// 		case "-":
// 			result = req.Operand1 - req.Operand2
// 		case "*":
// 			result = req.Operand1 * req.Operand2
// 		case "/":
// 			if req.Operand2 != 0 {
// 				result = req.Operand1 / req.Operand2
// 			} else {
// 				fmt.Println("Ошибка: Деление на ноль!")
// 				result = 0
// 			}
// 		default:
// 			fmt.Println("Ошибка: Неверная операция!")
// 			result = 0
// 		}
// 		req.Result <- result
// 	}
// }

// func client(operand1, operand2 float64, operation string, requests chan CalcRequest, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	resultChan := make(chan float64)

// 	request := CalcRequest{Operand1: operand1, Operand2: operand2, Operation: operation, Result: resultChan}
// 	requests <- request

// 	result := <-resultChan
// 	fmt.Printf("Результат операции %f %s %f = %f\n", operand1, operation, operand2, result)
// }

// func getUserInput() (float64, float64, string) {
// 	var op1, op2 float64
// 	var operation string

// 	fmt.Print("Введите первое число: ")
// 	fmt.Scan(&op1)

// 	fmt.Print("Введите второе число: ")
// 	fmt.Scan(&op2)

// 	fmt.Print("Введите операцию (+, -, *, /): ")
// 	fmt.Scan(&operation)

// 	return op1, op2, operation
// }

// func main() {
// 	requests := make(chan CalcRequest)
// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	go calculator(requests, &wg)

// 	for {
// 		var choice string

// 		op1, op2, operation := getUserInput()

// 		wg.Add(1)
// 		go client(op1, op2, operation, requests, &wg)

// 		time.Sleep(time.Millisecond * 500)
// 		fmt.Print("Хотите продолжить (да/нет)? ")
// 		fmt.Scan(&choice)
// 		if choice != "да" {
// 			break
// 		}
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(requests)
// 	}()

// 	wg.Wait()
// 	fmt.Println("Все операции завершены.")
// }
