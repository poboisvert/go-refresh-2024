package questions

import (
	"encoding/json"
	"fmt"
	"os"
)

type Question struct {
	Country string `json:"country"`
	Capital string `json:"capital"` // Corrected field name from "City" to "Capital" to match the JSON file
}

func LoadQuestions() ([]Question, error) {
	jsonData, err := os.ReadFile("quiz.json") // Removed extra parentheses around the file name
	if err != nil {
		return nil, fmt.Errorf("Error loading file: %w", err)
	}

	var questions []Question

	err = json.Unmarshal(jsonData, &questions)
	if err != nil {
		return nil, fmt.Errorf("Error reading JSON: %w", err)
	}

	return questions, nil
}
