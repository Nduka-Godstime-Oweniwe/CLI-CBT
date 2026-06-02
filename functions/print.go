package cbt

import (
	"fmt"
	"strings"
	"time"
)

func PrintQuestion(questions []Question, numb int) ([]string, int) {
	var userAnswers []string
	var answer string
	var score int
	for i, q := range questions {

		ClearScreen()
		fmt.Println("===> " + "Test Mode!" + " <===")
		fmt.Printf(
			"%d. %s\n",
			i+1,
			q.Question,
		)

		fmt.Printf(
			"A. %s\n",
			q.OptionA,
		)

		fmt.Printf(
			"B. %s\n\n",
			q.OptionB,
		)

		for {
			answer = strings.ToUpper(
				strings.TrimSpace(
					UserInput("Answer (A/B): "),
				),
			)
			if answer != "A" && answer != "B" {
				fmt.Println("Invalid Answer!")
				continue
			}
			break
		}

		userAnswers = append(userAnswers, answer)
		if answer == q.Answer {
			score++
		}
		// This prints the question when user is taking a
		// It takes a slice of strings and the index of the particular question to be printed
		// return user

	}
	return userAnswers, score
}

func PrintResults(userAnswers []string, score int, selectedSubjects []string, testQuestions []Question, numb int) {
	ClearScreen()
	fmt.Println("===> " + "Test Results!" + " <===")
	percent := (score * 100) / len(testQuestions)
	testScore := fmt.Sprintf("Score: %d/%d (%d%%)", score, len(testQuestions), percent)
	scores := LoadScore()
	if scores == nil {
		scores = make(map[string]Score)
	}
	scores[time.Now().Format("2006-01-02 15:04:05.000")] = Score{

		Subject: strings.Join(selectedSubjects, ", "),
		Correct: score,
		Total:   len(testQuestions),
		Score:   percent,
	}
	DumpScore(scores)
	fmt.Print(testScore + "\n")

	if percent >= 70 {
		fmt.Println("Congratulations! You passed the test.")
	} else {
		fmt.Println("Unfortunately, you did not pass the test. Better luck next time!")
	}

	if percent != 100 {
		PrintCorrectSolution(testQuestions, userAnswers)
		fmt.Println("\nReview the correct answers above and try again to improve your score!")
		fmt.Println("Would you like to retake this test?")
		if UserOption("1. Yes\n2. No\nSelect An Option: ", 2) == 1 {
			userAnswers, score = PrintQuestion(testQuestions, numb)
			PrintResults(userAnswers, score, selectedSubjects, testQuestions, numb)
		}
	}
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
