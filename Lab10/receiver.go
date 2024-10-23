package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
)

func loadPublicKey() *rsa.PublicKey {
	publicFile, _ := os.ReadFile("public_key_sender.pem")
	block, _ := pem.Decode(publicFile)
	publicKey, _ := x509.ParsePKIXPublicKey(block.Bytes)
	return publicKey.(*rsa.PublicKey)
}

func verifySignature(publicKey *rsa.PublicKey, message, signature string) bool {
	hash := sha256.Sum256([]byte(message))
	signatureBytes, _ := base64.StdEncoding.DecodeString(signature)
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signatureBytes)
	return err == nil
}

func main() {
	fmt.Println("Получатель загружает сообщение и подпись...")

	message, _ := os.ReadFile("message.txt")
	signature, _ := os.ReadFile("signature.txt")

	publicKey := loadPublicKey()

	if verifySignature(publicKey, string(message), string(signature)) {
		fmt.Println("Подпись действительна!")
		fmt.Println("Сообщение:", string(message))
	} else {
		fmt.Println("Подпись недействительна!")
	}
}
