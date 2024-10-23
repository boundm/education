package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"os"
	"strings"
)

func main() {
	fmt.Println("Выберите действие:")
	fmt.Println("1: Вычислить хэш")
	fmt.Println("2: Проверить целостность данных")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		solveHash()
	case 2:
		verifyHash()
	default:
		fmt.Println("Неверный выбор")
	}
}

func solveHash() {
	fmt.Println("Введите строку для хэширования:")
	var input string
	fmt.Scanln(&input)

	fmt.Println("Выберите алгоритм хэширования (md5, sha256, sha512):")
	var algo string
	fmt.Scanln(&algo)

	var hasher hash.Hash
	switch strings.ToLower(algo) {
	case "md5":
		hasher = md5.New()
	case "sha256":
		hasher = sha256.New()
	case "sha512":
		hasher = sha512.New()
	default:
		fmt.Println("Неверный алгоритм")
		os.Exit(1)
	}

	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	fmt.Printf("Хэш строки (%s): %s\n", algo, hashString)
}

func verifyHash() {
	fmt.Println("Введите строку:")
	var input string
	fmt.Scanln(&input)

	fmt.Println("Введите ожидаемый хэш:")
	var expectedHash string
	fmt.Scanln(&expectedHash)

	fmt.Println("Выберите алгоритм хэширования (md5, sha256, sha512):")
	var algo string
	fmt.Scanln(&algo)

	var hasher hash.Hash
	switch strings.ToLower(algo) {
	case "md5":
		hasher = md5.New()
	case "sha256":
		hasher = sha256.New()
	case "sha512":
		hasher = sha512.New()
	default:
		fmt.Println("Неверный алгоритм")
		os.Exit(1)
	}

	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)

	if hashString == expectedHash {
		fmt.Println("Хэши совпадают! Целостность данных подтверждена.")
	} else {
		fmt.Println("Хэши не совпадают. Целостность данных нарушена.")
	}
}
