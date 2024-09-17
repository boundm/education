package main

import "fmt"

func main() {
	str := "Hi there!"
	length := stringLength(str)
	fmt.Println("Длина строки:", length, "символов")
}

func stringLength(s string) int {
	return len(s)
}
