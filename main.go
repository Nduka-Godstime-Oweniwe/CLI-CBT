package main

import (
	cbt "CBT/functions"
	"fmt"
	"strings"
	"time"
)

func Subjects(mp map[string][]cbt.Question) []string {
	result := []string{}
	for key := range mp {
		result = append(result, key)
	}

	return result
}

// func IsContained(str string, slice []string) bool {
// 	for i := 0; i < len(slice); i++ {
// 		if slice[i] == str {
// 			return true
// 		}

// 	}
// 	return false
// }

// func RemoveIndex(i int, slice []cbt.Question) []string {
// 	result := []string{}
// 	for k := 0; k < len(slice); k++ {
// 		if k != i {
// 			result = append(result, slice[k])
// 		}

// 	}
// 	return result
// }

func PrintSubject(subject []string) {
	for i := 0; i < len(subject); i++ {
		fmt.Printf("%d. %s\n", i+1, subject[i])
	}
}

func main() {
	for {
		cbt.ClearScreen()
		fmt.Println("1. Take A Test")
		fmt.Println("2. Upload Questions")
		fmt.Println("3. Exit")
		option := cbt.UserOption("Select an option: ", 3)

		// Take A Test
		if option == 1 {
			questionbank := cbt.LoadQuestion()
			if len(questionbank) == 0 {
				fmt.Println("Oops! There are no available questions pls try again some other time")
				time.Sleep(2 * time.Second)
				continue
			}

			subjects := Subjects(questionbank)

			selectedSubjects := []string{}
			for {
				cbt.ClearScreen()
				PrintSubject(subjects)
				option := cbt.UserOption("Select An Option: ", len(subjects))
				if option != len(subjects)+1 {
					selectedSubjects = append(selectedSubjects, subjects[option-1])
					// subjects = RemoveIndex(option-1, subjects)
				}
				if len(subjects) == 0 {
					continue
				}
				exit := cbt.UserOption("1. Select Another Subject\n2. Proceed to Test\nSelect An Option: ", 2)
				time.Sleep(1 * time.Second)
				if exit == 2 {
					break
				}
			}

			// Start Test

			cbt.ClearScreen()

			CbtQuestions := []cbt.Question{}
			// numb := 0
			numb := cbt.UserOption("Enter Number Of Questions: ", 50)
			testQuestions := cbt.RandomQuestion(selectedSubjects, numb)
			if testQuestions == nil {
				fmt.Println("Oops! There are no available questions for the selected subjects pls try again some other time")
				time.Sleep(2 * time.Second)
				continue
			}
			if len(testQuestions) == 1 {
				fmt.Println(testQuestions[0])
				time.Sleep(2 * time.Second)
				continue
			}

			score := 0
			userAnswers := []string{}
			for i, q := range testQuestions {

				cbt.ClearScreen()
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

				answer := strings.ToUpper(
					strings.TrimSpace(
						cbt.UserInput("Answer (A/B): "),
					),
				)
				userAnswers = append(userAnswers, answer)
				if answer == q.Answer {
					score++
				}
			}

			// Results
			cbt.ClearScreen()

			fmt.Printf("Your Score: %v%%\n\n", (score*100)/(len(testQuestions)/2))
			cbt.PrintCorrectSolution(CbtQuestions, userAnswers)
			cbt.UserInput("Press anything to continue")
		}

		// Upload Questions
		if option == 2 {
			cbt.ClearScreen()

			mp := cbt.LoadQuestion()
			subjects := Subjects(mp)
			PrintSubject(subjects)
			fmt.Printf("%v. Upload New Subject\n", len(subjects)+1)
			fmt.Printf("%v. Go back\n", len(subjects)+2)
			option := cbt.UserOption("Select An Option: ", len(subjects)+2)
			if option == len(subjects)+2 {
				// This is when user wants to go back to previous menu
				continue
			}
			if option == len(subjects)+1 {
				nameOfNewSubject := cbt.UserInput("Enter Subject Name: ")
				subjects = append(subjects, nameOfNewSubject)
				cbt.Upload(nameOfNewSubject, mp)
			} else {
				cbt.Upload(subjects[option-1], mp)
			}

			cbt.UserInput("Press Any key to continue")
			continue

		}
		// Exit
		if option == 3 {
			break
		}
	}

}
