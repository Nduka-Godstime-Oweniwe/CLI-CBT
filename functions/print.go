package cbt

import (
	"fmt"
	"strings"
)

func PrintQuestion(questions []Question, numb int) string {
	fmt.Printf("%s\n", questions[numb].Question)
	fmt.Printf("A. %s\n", questions[numb].OptionA)
	fmt.Printf("B. %s\n\n", questions[numb].OptionB)

	answer := ""
	for {
		answer = UserInput("Enter Answer(A or B)")
		answer = strings.TrimSpace(strings.ToUpper(answer))
		if answer != "A" && answer != "B" {
			fmt.Println("Invalid Answer!")
		} else {
			break
		}
	}

	// This prints the question when user is taking a cbt.
	// It takes a slice of strings and the index of the particular question to be printed
	// return user
	return answer
}

func PrintCorrectSolution(questions []Question, userAnswers []string) {
	// After User is done with the test. All questions asked, their answers, the user's answer would be appended into a slice
	// E.g {"What is my name\nA. Nduka\nB. Joshua","A","C"}
	// This means index 0 was the question, index 1 was the correct answers, index 2 was the user answer
	// Hence the questions have indexes 0,3,6,9 and so on...
	//The correct answers have indexes 1,4,7,10 and so on...
	//The user answer are indexes 2,5,8,11 and so on...
	// You're to print the question, Correct Answer,  user Answer. Example Bellow
	// What is my Name
	// A. Nduka
	// B. Joshua
	// Correct Answer : A
	// Your Answer : C
	for i := 0; i < len(questions); i++ {

		fmt.Printf(
			"%d. %s\n",
			i+1,
			questions[i].Question,
		)

		fmt.Printf(
			"A. %s\n",
			questions[i].OptionA,
		)

		fmt.Printf(
			"B. %s\n\n",
			questions[i].OptionB,
		)

		fmt.Printf(
			"Correct Ans: %v\n",
			questions[i].Answer,
		)

		fmt.Printf(
			"Your Answer: %v\n",
			userAnswers[i],
		)
	}
}
