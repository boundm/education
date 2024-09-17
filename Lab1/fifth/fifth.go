package main

import "fmt"

func sum(a, b float64) float64 {
	return a + b
}

func difference(a, b float64) float64 {
	return a - b
}

func main() {
	num1 := 11.3
	num2 := 4.5

	s := sum(num1, num2)
	d := difference(num1, num2)

	fmt.Println("Первое число:", num1)
	fmt.Println("Второе число:", num2)
	fmt.Println("Сумма:", s)
	fmt.Println("Разность:", d)
}
