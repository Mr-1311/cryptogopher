package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

var encryption = flag.BoolP("encryption", "e", false, "encryption flag")
var decryption = flag.BoolP("decryption", "d", false, "decryption flag")

func main() {

	flag.Parse()

	if *encryption == false && *decryption == false {
		fmt.Println("Chose the main operation: [-e] or [-d]")
		return
	}

	if *encryption == true && *decryption == true {
		fmt.Println("Main operations can't set both at same time': [-e] and [-d]")
		return
	}

	args := os.Args[1:]
	path := args[len(args)-1]

	if *encryption == true {
		CryptoHandler(path, 0)
	}

	if *decryption == true {
		CryptoHandler(path, 1)
	}
}
