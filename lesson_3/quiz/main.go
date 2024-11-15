package main

import (
	"fmt"
	"os"
	data "quiz/data"
)

func main() {
	questions, err := data.LoadQuestions()
	if err != nil {
		fmt.Printf("Error loading file")
		os.Exit(1)
	}
}
