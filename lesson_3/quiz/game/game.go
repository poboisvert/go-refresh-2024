package game

import (
	"bufio"
	"fmt"
	"os"
	data "quiz/data"
	"strings"
)

func Run(questions []data.Question) uint {
	fmt.Println("Country Quiz Started!")
	var correctAnswers uint
	for _, question := range questions {
		if askQuestion(question) {
			correctAnswers++
		}
	}
	return correctAnswers
}

func askQuestion(questions data.Question) bool {
	fmt.Printf("\n Enter the Capital of: %s", questions.Country)

	if getUserInput() == strings.ToLower(questions.Capital) {
		fmt.Printf("Correct answer")
		return true
	} else {
		fmt.Printf("Correct answer. Answer is: ", questions.Capital)
		return false
	}
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Your Answer: ")
		result, err := reader.ReadString('\n')

		if err != nil {
			fmt.Printf("Error User Input: %v\n", err)
			continue
		}

		return strings.ToLower(strings.TrimRight(result, "\r\n"))
	}
}
