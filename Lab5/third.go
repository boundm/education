package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите радиус круга: ")
	scanner.Scan()
	input := scanner.Text()

	radius, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Ошибка: неверный формат радиуса")
		return
	}

	circle := Circle{radius: radius}

	area := circle.Area()
	fmt.Printf("Площадь круга с радиусом %.2f: %.2f\n", circle.radius, area)
}
