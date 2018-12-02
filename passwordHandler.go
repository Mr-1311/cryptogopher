package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

func PasswordHandler(operation int) string {
	fmt.Print("\nPlease enter password: ")
	password, _ := terminal.ReadPassword(int(os.Stdin.Fd()))
	if operation == 0 {
		fmt.Print("\nPlease enter password (again): ")
		password2, _ := terminal.ReadPassword(int(os.Stdin.Fd()))
		if strings.Compare(string(password), string(password2)) != 0 {
			fmt.Println("\nPasswords do not match, please type again!")
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
