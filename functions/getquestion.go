package cbt

import (
	"fmt"
	"strings"
)

func UserQuestion() (string, string, string, string) {
	var question string
	var optionA string
	var optionB string
	var answer string
	// This function Uses UserInput("Enter Question ")

	for {
		question = UserInput("Enter Question: ")
		if question == "" {
			fmt.Print("Question cannot be empty\n")
			continue
		}
		break
	}
	//UserInput("Enter Option A: ")
	for {
		optionA = UserInput("Enter Option A: ")
		if optionA == "" {
			fmt.Print("Options cannot be empty\n")
			continue
		}
		break
	}
	// UserInput("Enter Option B: ")
	for {
		optionB = UserInput("Enter Option B: ")
		if optionB == "" {
			fmt.Print("Options cannot be empty\n")
			continue
		}
		break
	}
	// UserInput("Enter Answer : ")
	// Answers must be A or B. If user types in a or b convert to upper.
	for {
		answer = UserInput("Enter Answer (A or B): ")
		answer = strings.ToUpper(answer)
		if answer == "" {
			fmt.Print("Answer cannot be empty\n")
			continue
		}
		if answer != "A" && answer != "B" {
			fmt.Print("Answer must be A or B\n")
			continue
		}
		break
	}

	//Return both values
	return question, optionA, optionB, answer

}
