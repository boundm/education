package main

import "fmt"

func main() {
	num1 := 5
	num2 := 10
	average := calculateAverage(num1, num2)
	fmt.Println("Среднее значение:", average)
}

func calculateAverage(a int, b int) float64 {
	return float64(a+b) / 2.0
}
