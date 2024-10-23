package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
)

func generateKeys() {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)

	// Сохранение закрытого ключа
	privateFile, _ := os.Create("private_key_sender.pem")
	defer privateFile.Close()
	privatePem := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	pem.Encode(privateFile, privatePem)

	publicFile, _ := os.Create("public_key_sender.pem")
	defer publicFile.Close()
	publicKeyBytes, _ := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	publicPem := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	pem.Encode(publicFile, publicPem)

	fmt.Println("Ключи отправителя успешно сгенерированы и сохранены.")
}

func signMessage(privateKey *rsa.PrivateKey, message string) string {
	hash := sha256.Sum256([]byte(message))
	signature, _ := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	return base64.StdEncoding.EncodeToString(signature)
}

func loadPrivateKey() *rsa.PrivateKey {
	privateFile, _ := os.ReadFile("private_key_sender.pem")
	block, _ := pem.Decode(privateFile)
	privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return privateKey
}

func main() {
	generateKeys()

	fmt.Println("Введите сообщение для подписи:")
	var message string
	fmt.Scanln(&message)

	privateKey := loadPrivateKey()
	signature := signMessage(privateKey, message)

	os.WriteFile("message.txt", []byte(message), 0644)
	os.WriteFile("signature.txt", []byte(signature), 0644)

	fmt.Println("Сообщение и подпись отправлены получателю (сохранены в файлы message.txt и signature.txt)")
}
