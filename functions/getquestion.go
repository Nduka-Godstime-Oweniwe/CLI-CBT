package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func UserInput(prompt string) string {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print(prompt)

	input, _ := reader.ReadString('\n')

	// Remove newline character
	input = strings.TrimSpace(input)

	return input
}

func UserQuestion() (string, string) {
	var question string
	var optionA string
	var optionB string
	var answer string
	// This function Uses UserInput("Enter Question ")
	for {

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
		for {
			answer = UserInput("Enter Answer (A or B): ")
			if answer == "" {
				fmt.Print("Answer cannot be empty\n")
				continue
			}
			break
		}

		// Answers must be A or B. If user types in a or b convert to upper.
		for {
			answer = strings.ToUpper(answer)
			if answer != "A" && answer != "B" {
				fmt.Print("Answer must be A or B\n")
				continue
			}
			break
		}
		break
	}

	// Joint Question, OptionA and OptionB together as one string but separated by a newline character "\n"
	fullQuestion := question +
		"\nA. " + optionA +
		"\nB. " + optionB

	//Return both values
	return fullQuestion, answer

}

func main() {
	UserQuestion()
}
