package main

import "fmt"

func main() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scanln(&number)
	fmt.Println(checker(number))
}

func checker(number int) string {
	if number == 0 {
		return "Zero"
	}
	if number > 0 {
		return "Positive"
	}
	if number < 0 {
		return "Negative"
	}
	return "Ошибка"

}
