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

	// Инициализация переменной для хранения суммы
	sum := 0

	// Преобразование строк в числа и подсчет суммы
	for _, numStr := range numbersStr {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Printf("Ошибка: неверный формат числа '%s'\n", numStr)
			return
		}
		sum += num
	}

	// Вывод суммы
	fmt.Printf("Сумма чисел: %d\n", sum)
}
