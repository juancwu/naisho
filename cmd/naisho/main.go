package main

import (
	"fmt"
	"log"
	"os"
    "encoding/hex"

    "github.com/juancwu/naisho/internal/crypto"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    keyBytes, err := hex.DecodeString(os.Getenv("SECRET_KEY"))
    if err != nil {
        panic(err.Error())
    }

    fmt.Println("Decoded HEX:", keyBytes)
    fmt.Println("Decoded key length:", len(keyBytes))

    plaintext := []byte("some content")

    aesgcm, err := crypto.New(keyBytes)
    if err != nil {
        panic(err.Error())
    }

    // encrypt data
    nonce := make([]byte, aesgcm.NonceSize())
    ciphertext, nonce, err := crypto.Encrypt(aesgcm, plaintext)
    if err != nil {
        panic(err.Error())
    }

    fmt.Printf("Nonce: %x\n", nonce)
    fmt.Printf("Ciphertext: %x\n", ciphertext)

    // decrypt the data
    decrypted, err := crypto.Decrypt(aesgcm, ciphertext, nonce) 
    if err != nil {
        panic(err.Error())
    }

    fmt.Printf("Decrypted Text: %s\n", decrypted)
}
