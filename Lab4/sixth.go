package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Создание сканера для чтения ввода пользователя
	scanner := bufio.NewScanner(os.Stdin)

	// Запрос строки чисел у пользователя
	fmt.Print("Введите несколько чисел, разделенных пробелами: ")
	scanner.Scan()
	input := scanner.Text()

	// Разделение строки на отдельные числа
	numbersStr := strings.Fields(input)

	// Инициализация среза для хранения чисел
	numbers := make([]int, len(numbersStr))

	// Преобразование строк в числа
	for i, numStr := range numbersStr {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Printf("Ошибка: неверный формат числа '%s'\n", numStr)
			return
		}
		numbers[i] = num
	}

	// Вывод чисел в обратном порядке
	fmt.Println("Числа в обратном порядке:")
	for i := len(numbers) - 1; i >= 0; i-- {
		fmt.Print(numbers[i], " ")
	}
	fmt.Println()
}
