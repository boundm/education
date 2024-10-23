package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Выберите действие:")
	fmt.Println("1: Зашифровать строку")
	fmt.Println("2: Расшифровать строку")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		encrypt()
	case 2:
		decrypt()
	default:
		fmt.Println("Неверный выбор")
	}
}

func encrypt() {
	fmt.Println("Введите строку для шифрования:")
	var plaintext string
	fmt.Scanln(&plaintext)

	fmt.Println("Введите секретный ключ (32 байта для AES-256):")
	var key string
	fmt.Scanln(&key)

	keyBytes := []byte(key)
	if len(keyBytes) != 32 {
		fmt.Println("Ключ должен быть длиной 32 байта (256 бит)")
		os.Exit(1)
	}

	ciphertext, err := aesEncrypt([]byte(plaintext), keyBytes)
	if err != nil {
		fmt.Println("Ошибка при шифровании:", err)
		os.Exit(1)
	}

	fmt.Println("Зашифрованная строка:", base64.StdEncoding.EncodeToString(ciphertext))
}

func decrypt() {
	fmt.Println("Введите зашифрованную строку (Base64):")
	var encryptedBase64 string
	fmt.Scanln(&encryptedBase64)

	fmt.Println("Введите секретный ключ (32 байта для AES-256):")
	var key string
	fmt.Scanln(&key)

	keyBytes := []byte(key)
	if len(keyBytes) != 32 {
		fmt.Println("Ключ должен быть длиной 32 байта (256 бит)")
		os.Exit(1)
	}

	ciphertext, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		fmt.Println("Ошибка при декодировании Base64:", err)
		os.Exit(1)
	}

	plaintext, err := aesDecrypt(ciphertext, keyBytes)
	if err != nil {
		fmt.Println("Ошибка при расшифровании:", err)
		os.Exit(1)
	}

	fmt.Println("Расшифрованная строка:", string(plaintext))
}

func aesEncrypt(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

func aesDecrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("длина зашифрованных данных меньше минимального размера nonce")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
