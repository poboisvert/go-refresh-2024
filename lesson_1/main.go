package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {
	fmt.Println("hello")

	var x int = 1
	fmt.Printf("number: %d\n", x)

	file, _ := os.Create("output.txt")
	fmt.Fprintf(file, "number: %d\n", x)
	file.Close()

	fmt.Println("Enter hex number")
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)

	input = strings.ToLower(input)
	if input == "stop" {
		fmt.Println("You typed stop")
	}

	i := new(big.Int)
	if _, err := i.SetString(processHex(input), 16); !err {
		// 0xbc456 -> bc456
		fmt.Println("Invalid HEX")
	}
	fmt.Println(i)

}

func processHex(hexStr string) string {
	return strings.TrimPrefix(hexStr, "0x")
}
