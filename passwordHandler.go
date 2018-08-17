package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func PasswordHandler(operation int) string {
	fmt.Print("enter password:")
	password, _ := terminal.ReadPassword(int(os.Stdin.Fd()))
	if operation == 0 {
		fmt.Print("enter password again:")
		password2, _ := terminal.ReadPassword(int(os.Stdin.Fd()))
		if string(password) != string(password2) {
			fmt.Print("passwords are not match!")
			return PasswordHandler(operation)
		}
	}

	hashPass := CreateHash(password)
	return hashPass
}

func CreateHash(key []byte) string {
	hasher := md5.New()
	hasher.Write(key)
	return hex.EncodeToString(hasher.Sum(nil))
}
