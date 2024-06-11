package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"bytes"

	"golang.org/x/crypto/pbkdf2"
)

func encryptFile(password, salt []byte, filepath string) error {
	key := pbkdf2.Key(password, salt, 4096, 32, sha256.New)

	plaintext, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	// Add padding
	padding := aes.BlockSize - len(plaintext)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	plaintext = append(plaintext, padtext...)

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	copy(ciphertext[:aes.BlockSize], iv)

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	encodedCiphertext := base64.StdEncoding.EncodeToString(ciphertext)

	return ioutil.WriteFile(filepath+".enc", []byte(encodedCiphertext), 0644)
}


func decryptFile(password, salt []byte, filepath string) error {
	key := pbkdf2.Key(password, salt, 4096, 32, sha256.New)

	ciphertext, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	decodedCiphertext, err := base64.StdEncoding.DecodeString(string(ciphertext))
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	if len(decodedCiphertext) < aes.BlockSize {
		return fmt.Errorf("ciphertext too short")
	}
	iv := decodedCiphertext[:aes.BlockSize]
	decodedCiphertext = decodedCiphertext[aes.BlockSize:]

	if len(decodedCiphertext)%aes.BlockSize != 0 {
		return fmt.Errorf("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decodedCiphertext, decodedCiphertext)

	// Remove padding
	padding := decodedCiphertext[len(decodedCiphertext)-1]
	decodedCiphertext = decodedCiphertext[:len(decodedCiphertext)-int(padding)]

	return ioutil.WriteFile(filepath+".dec", decodedCiphertext, 0644)
}


func main() {
	// Replace with your desired password
	password := []byte("your_secret_password")
	salt := []byte("your_salt")

	// Specify the file to encrypt/decrypt
	filepath := "your_file.txt"

	// Choose between encryption or decryption
	action := "decrypt" // or "decrypt"

	if action == "encrypt" {
		err := encryptFile(password, salt, filepath)
		if err != nil {
			fmt.Println("Error encrypting file:", err)
			return
		}
		fmt.Println("File encrypted successfully!")
	} else if action == "decrypt" {
		err := decryptFile(password, salt, filepath+".enc")
		if err != nil {
			fmt.Println("Error decrypting file:", err)
			return
		}
		fmt.Println("File decrypted successfully!")
	} else {
		fmt.Println("Invalid action. Please choose 'encrypt' or 'decrypt'.")
	}
}