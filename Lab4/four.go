package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Создание сканера для чтения ввода пользователя
	scanner := bufio.NewScanner(os.Stdin)

	// Запрос строки у пользователя
	fmt.Print("Введите строку: ")
	scanner.Scan()
	input := scanner.Text()

	// Преобразование строки в верхний регистр
	upperCaseInput := strings.ToUpper(input)

	// Вывод строки в верхнем регистре
	fmt.Println("Строка в верхнем регистре:", upperCaseInput)
}
