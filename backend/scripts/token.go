package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

func GenerateSecretKey(length int) string {
    key := make([]byte, length)
    _, err := rand.Read(key)
    if err != nil {
        log.Fatalf("Failed to generate random key: %v", err)
    }
    return base64.URLEncoding.EncodeToString(key)
}

func main() {
    secretKey := GenerateSecretKey(32) 
    fmt.Println("New Secret Key:", secretKey)
}