package main

import "fmt"

func main() {

	num1 := 15
	num2 := 5

	sum := num1 + num2
	difference := num1 - num2
	product := num1 * num2
	quotient := num1 / num2
	remainder := num1 % num2

	fmt.Println("Первое число:", num1)
	fmt.Println("Второе число:", num2)
	fmt.Println("Сумма:", sum)
	fmt.Println("Разность:", difference)
	fmt.Println("Произведение:", product)
	fmt.Println("Частное:", quotient)
	fmt.Println("Остаток от деления:", remainder)
}
