// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// type Person struct {
// 	name string
// 	age  int
// }

// func (p Person) DisplayInfo() {
// 	fmt.Printf("Имя: %s, Возраст: %d лет\n", p.name, p.age)
// }

// func (p *Person) Birthday() {
// 	p.age++
// }

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	person1 := Person{name: "Pasha", age: 20}
// 	person2 := Person{name: "Dima", age: 100}

// 	person1.DisplayInfo()
// 	person2.DisplayInfo()

// 	fmt.Print("Хотите увеличить возраст Pasha на 1 год? (да/нет): ")
// 	scanner.Scan()
// 	input := scanner.Text()
// 	if strings.ToLower(input) == "да" {
// 		person1.Birthday()
// 	}

// 	fmt.Print("Хотите увеличить возраст Dima на 1 год? (да/нет): ")
// 	scanner.Scan()
// 	input = scanner.Text()
// 	if strings.ToLower(input) == "да" {
// 		person2.Birthday()
// 	}
// 	person1.DisplayInfo()
// 	person2.DisplayInfo()
// }
