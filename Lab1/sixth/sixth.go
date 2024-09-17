package main

import "fmt"

func average(a, b, c float64) float64 {
	return (a + b + c) / 3
}

func main() {
	num1 := 18.34
	num2 := 2.6
	num3 := 30.1

	avg := average(num1, num2, num3)

	fmt.Println("Среднее значение чисел", num1, num2, num3, "равно:", avg)
}
