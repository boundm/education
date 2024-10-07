// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"sync"
// )

// type Task struct {
// 	Text   string
// 	Result chan string
// }

// func worker(id int, tasks chan Task, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for task := range tasks {
// 		reversed := reverseString(task.Text)
// 		fmt.Printf("Воркер %d обработал строку: %s\n", id, reversed)
// 		task.Result <- reversed
// 	}
// }

// func reverseString(s string) string {
// 	runes := []rune(s)
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	return string(runes)
// }

// func readLines(filename string) ([]string, error) {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	var lines []string
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		lines = append(lines, scanner.Text())
// 	}
// 	return lines, scanner.Err()
// }

// func writeResultsToFile(filename string, results []string) error {
// 	file, err := os.Create(filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	writer := bufio.NewWriter(file)
// 	for _, result := range results {
// 		_, err := writer.WriteString(result + "\n")
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return writer.Flush()
// }

// func main() {
// 	var numWorkers int

// 	fmt.Print("Введите количество воркеров: ")
// 	fmt.Scan(&numWorkers)

// 	lines, err := readLines("input.txt")
// 	if err != nil {
// 		fmt.Println("Ошибка чтения файла:", err)
// 		return
// 	}

// 	tasks := make(chan Task, len(lines))
// 	results := make(chan string, len(lines))

// 	var wg sync.WaitGroup

// 	for i := 1; i <= numWorkers; i++ {
// 		wg.Add(1)
// 		go worker(i, tasks, &wg)
// 	}

// 	for _, line := range lines {
// 		task := Task{Text: line, Result: results}
// 		tasks <- task
// 	}

// 	close(tasks)

// 	wg.Wait()

// 	var reversedLines []string
// 	for i := 0; i < len(lines); i++ {
// 		reversedLines = append(reversedLines, <-results)
// 	}

// 	close(results)

// 	err = writeResultsToFile("output.txt", reversedLines)
// 	if err != nil {
// 		fmt.Println("Ошибка записи в файл:", err)
// 		return
// 	}

// 	fmt.Println("Реверсирование строк завершено. Результаты записаны в файл output.txt.")
// }
