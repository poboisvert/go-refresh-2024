package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"xor/cipherer"
)

var cipherMode = flag.String("mode", "cipher", "Enable cipher mode, This is by default")
var secretKey = flag.String("secret", "", "Secretkey must have 1 character")

// go run . --cipher --secret key
func main() {
	flag.Parse()

	switch *cipherMode {
	case "cipher":
		plaintext := getUserInput("Enter your text to cipher: ")
		fmt.Println(plaintext)
		cipheredText, err := cipherer.Cipher(plaintext, *secretKey)
		if err != nil {
			fmt.Println("Error during ciphering:", err)
		} else {
			fmt.Println(cipheredText)
		}

	case "decipher":
		cipheredText := getUserInput("Enter your text to decipher: ")
		fmt.Println(cipheredText)
		decipheredText, err := cipherer.Decipher(cipheredText, *secretKey)
		if err != nil {
			fmt.Println("Error during deciphering:", err)
		} else {
			fmt.Println(decipheredText)
		}
	default:
		fmt.Println("invalid Mode. Please use 'cipher' or 'decipher'")
		os.Exit(1)
	}

}

func getUserInput(msg string) string {
	fmt.Print(msg)

	reader := bufio.NewReader(os.Stdin)
	result, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred while reading text")
	}

	return strings.TrimRight(result, "\n")
}
