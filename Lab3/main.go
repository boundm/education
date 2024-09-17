package main

import (
	"Lab3/mathutils"
	"Lab3/stringutils"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите число для вычисления факториала: ")
	scanner.Scan()
	var number int
	fmt.Sscan(scanner.Text(), &number)

	factorial := mathutils.Factorial(number)
	fmt.Printf("Факториал числа %d: %d\n", number, factorial)

	fmt.Print("Введите строку для переворота: ")
	scanner.Scan()
	inputString := scanner.Text()

	reversedString := stringutils.Reverse(inputString)
	fmt.Printf("Перевернутая строка: %s\n", reversedString)

	fmt.Println("Запускаем генератор массива?\nДа - 1\nНет - 2")
	scanner.Scan()
	var userAnswer int
	fmt.Sscan(scanner.Text(), &userAnswer)
	if userAnswer == 1 {
		array_generator()
	} else if userAnswer == 2 {
		fmt.Println("Программа завершена")
		os.Exit(0)
	} else {
		fmt.Println("Неверный ввод")
		os.Exit(1)
	}

	fmt.Println("Запускаем срез строк?\nДа - 1\nНет - 2")
	scanner.Scan()
	var userAnswer1 int
	fmt.Sscan(scanner.Text(), &userAnswer1)
	if userAnswer1 == 1 {
		string_cutter()
	} else if userAnswer1 == 2 {
		fmt.Println("Программа завершена")
		os.Exit(0)
	} else {
		fmt.Println("Неверный ввод")
		os.Exit(1)
	}
}

func array_generator() {
	rand.Seed(time.Now().UnixNano())

	var numbers [5]int

	for i := 0; i < 5; i++ {
		numbers[i] = rand.Intn(100)
	}

	fmt.Println("Значения массива:")
	for i, value := range numbers {
		fmt.Printf("numbers[%d] = %d\n", i, value)
	}

	slice := numbers[:]

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\nВыберите действие:\n1 - Добавить элемент\n2 - Удалить элемент")
	scanner.Scan()
	var action int
	fmt.Sscan(scanner.Text(), &action)

	if action == 1 {
		var value, index int
		fmt.Println("Введите значение для добавления:")
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &value)
		fmt.Println("Введите индекс для добавления (0 -", len(slice), "):")
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &index)

		if index < 0 || index > len(slice) {
			fmt.Println("Неверный индекс")
		} else {
			slice = append(slice[:index], append([]int{value}, slice[index:]...)...)
		}
	} else if action == 2 {
		var index int
		fmt.Println("Введите индекс для удаления (0 -", len(slice)-1, "):")
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &index)

		if index < 0 || index >= len(slice) {
			fmt.Println("Неверный индекс")
		} else {
			slice = append(slice[:index], slice[index+1:]...)
		}
	} else {
		fmt.Println("Неверное действие")
	}

	fmt.Println("\nЗначения среза после выполнения действия:")
	for i, value := range slice {
		fmt.Printf("slice[%d] = %d\n", i, value)
	}
}

func string_cutter() {
	var strings []string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите строки (пустая строка для завершения):")
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		strings = append(strings, line)
	}

	fmt.Println("\nЗначения среза:")
	for i, value := range strings {
		fmt.Printf("strings[%d] = %s\n", i, value)
	}

	var longestString string
	for _, value := range strings {
		if len(value) > len(longestString) {
			longestString = value
		}
	}

	fmt.Printf("\nСамая длинная строка: %s\n", longestString)
}
