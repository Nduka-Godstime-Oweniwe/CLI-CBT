package main

import (
	cbt "CBT/functions"
	"fmt"
	"strings"
	"time"
	"unicode"
)

func Subjects(mp map[string][]cbt.Question) []string {
	result := []string{}
	for key := range mp {
		result = append(result, key)
	}
	return result
}

func PrintSubject(subject []string) []string {
	for i := 0; i < len(subject); i++ {
		fmt.Printf("%d. %s\n", i+1, subject[i])
	}
	return subject
}

func main() {
	for {
		cbt.ClearScreen()
		fmt.Println("===> " + "Welcome to CLI CBT!" + " <===")
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
			option := 0
			for {
				cbt.ClearScreen()
				fmt.Println("===> " + "Select Subjects!" + " <===")
				PrintSubject(subjects)
				fmt.Printf("%v. Select Multiple Subjects\n", len(subjects)+1)
				fmt.Printf("%v. Select All Subjects\n", len(subjects)+2)
				fmt.Printf("%v. Go back\n", len(subjects)+3)
				option = cbt.UserOption("Select An Option: ", len(subjects)+3)
				if len(subjects) == 0 {
					continue
				}
				if option <= len(subjects) && option > 0 {
					selectedSubjects = append(selectedSubjects, subjects[option-1])
					exit := cbt.UserOption("1. Proceed to Test\n2. Select Different Subject\nSelect An Option: ", 2)
					time.Sleep(1 * time.Second)
					if exit == 1 {
						break
					} else {
						selectedSubjects = nil
						continue
					}

				}
				if option == len(subjects)+1 {
					for {
						nameOfSubject := strings.TrimSpace(
							cbt.UserInput("Enter Subject Name or done to finish: "),
						)
						runesOfSubject := []rune(nameOfSubject)
						runesOfSubject[0] = unicode.ToUpper(runesOfSubject[0])
						nameOfSubject = string(runesOfSubject)
						if strings.ToLower(nameOfSubject) == "done" {
							break
						}
						if _, exists := questionbank[nameOfSubject]; !exists {
							fmt.Println("Subject does not exist!")
							time.Sleep(1 * time.Second)
							continue
						}
						selectedSubjects = append(selectedSubjects, nameOfSubject)
					}
					break
				}
				if option == len(subjects)+2 {
					selectedSubjects = subjects
					break
				}
				if option == len(subjects)+3 {
					break
				}

			}
			if option == len(subjects)+3 {
				continue
			}

			// Start Test

			cbt.ClearScreen()
			fmt.Println("===> " + "Starting Test!" + " <===")
			fmt.Printf("Selected Subjects: %s\n", strings.Join(selectedSubjects, ", "))

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
			for {
				if len(testQuestions) < numb {
					fmt.Printf("Only %d questions are available for the selected subjects. Do you want the test to proceed with the available questions?\n", len(testQuestions))
					time.Sleep(1 * time.Second)
					if cbt.UserOption("1. Yes\n2. No\nSelect An Option: ", 2) == 1 {
						break
					} else {
						numb = cbt.UserOption("Enter Number Of Questions: ", 50)
					}
				}

				testQuestions = cbt.RandomQuestion(selectedSubjects, numb)
				break
			}

			score := 0
			userAnswers := []string{}
			userAnswers, score = cbt.PrintQuestion(testQuestions, numb)
			cbt.PrintResults(userAnswers, score, selectedSubjects, testQuestions, numb)

			// Results
			// continue
			cbt.UserInput("Press anything to continue")
			continue
		}

		// Upload Questions
		if option == 2 {
			cbt.ClearScreen()
			fmt.Println("===> " + "Upload Questions!" + " <===")

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
