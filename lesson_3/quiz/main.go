package main

import (
	"fmt"
	"os"
	data "quiz/data"
	"quiz/game"
	lib "quiz/lib"
)

func main() {
	questions, err := data.LoadQuestions()
	if err != nil {
		fmt.Printf("Error loading file")
		os.Exit(1)
	}

	lib.Shuffle(questions)

	correctAnswer := game.Run(questions)

	fmt.Printf("Correct answer: %d", correctAnswer)

}
