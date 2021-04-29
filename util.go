package main

import (
	"encoding/hex"
	"os"

	"golang.org/x/crypto/sha3"
)

func hashThis(stringToHash string, size int) string {
	hashedData := []byte(stringToHash)
	slot := make([]byte, size)
	sha3.ShakeSum256(slot, hashedData)
	return hex.EncodeToString(slot[:])
}

func createDirIfItDontExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			println("couldnt create dir", err)
		}
	}
}

func checkDirs() {
	createDirIfItDontExist("invoices")
	createDirIfItDontExist("sales")
}

func createFile(filename string) {
	var _, err = os.Stat(filename)
	if os.IsNotExist(err) {
		var file, err = os.Create(filename)
		if err != nil {
			println("couldnt create file", err)
		}
		defer file.Close()
	}
}
func writeFile(filename, textToWrite string) {
	var file, err = os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		println("couldnt create file", err)
	}
	defer file.Close()
	_, err = file.WriteString(textToWrite)
	if err != nil {
		println("couldnt create file", err)
	}
	err = file.Sync()
	if err != nil {
		println("couldnt create file", err)
	}
}
func createInvoice(name, content string) {
	createFile(name)
	writeFile(name, content)
}
