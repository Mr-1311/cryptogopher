package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var key string

// operation 0 = enc 1 = dec
func CryptoHandler(path string, operation int) {

	isDirectory := checkDirectory(path)

	if isDirectory == true {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range files {
			if isDirectory {
				CryptoHandler(path+"/"+f.Name(), operation)
			}
		}
	} else {
		encryptAndWrite(path, operation)
	}

}

func encryptAndWrite(path string, operation int) {

	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}

	if key == "" {
		key = PasswordHandler(operation)
	}

	var text []byte = nil
	if operation == 0 {
		text = Encrypt(b, []byte(key))
	}
	if operation == 1 {
		text = Decrypt(b, []byte(key))
	}
	ioutil.WriteFile(path, text, os.ModeAppend)
}

// checkDirectory ...
func checkDirectory(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
	}

	return stat.IsDir()
}
