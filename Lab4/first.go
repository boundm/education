// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	// Создание и инициализация карты
// 	ages := make(map[string]int)

// 	// Добавление начальных записей
// 	ages["Andrey"] = 30
// 	ages["Igor"] = 25
// 	ages["Pashok"] = 20

// 	// Создание сканера для чтения ввода пользователя
// 	scanner := bufio.NewScanner(os.Stdin)

// 	// Запрос имени и возраста у пользователя
// 	fmt.Print("Введите имя: ")
// 	scanner.Scan()
// 	name := scanner.Text()

// 	fmt.Print("Введите возраст: ")
// 	scanner.Scan()
// 	ageStr := scanner.Text()

// 	// Преобразование возраста из строки в целое число
// 	age, err := strconv.Atoi(ageStr)
// 	if err != nil {
// 		fmt.Println("Ошибка: неверный формат возраста")
// 		return
// 	}

// 	// Добавление нового человека в карту
// 	ages[name] = age

// 	// Вывод всех записей на экран
// 	fmt.Println("Имена и возрасты:")
// 	for name, age := range ages {
// 		fmt.Printf("%s: %d лет\n", name, age)
// 	}
// 	avgAge := averageAge(ages)
// 	fmt.Printf("Средний возраст: %.2f лет\n", avgAge)

// 	// Запрос имени для удаления у пользователя
// 	fmt.Print("Введите имя для удаления: ")
// 	scanner.Scan()
// 	nameToRemove := scanner.Text()

// 	// Удаление записи по имени
// 	removePerson(ages, nameToRemove)

// 	// Вывод записей после удаления
// 	fmt.Println("\nИмена и возрасты после удаления:")
// 	for name, age := range ages {
// 		fmt.Printf("%s: %d лет\n", name, age)
// 	}
// }

// func averageAge(ages map[string]int) float64 {
// 	totalAge := 0
// 	count := 0

// 	for _, age := range ages {
// 		totalAge += age
// 		count++
// 	}

// 	if count == 0 {
// 		return 0
// 	}

// 	return float64(totalAge) / float64(count)
// }

// func removePerson(ages map[string]int, name string) {
// 	delete(ages, name)
// }
