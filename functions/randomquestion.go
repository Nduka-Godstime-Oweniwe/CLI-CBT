package cbt

import (
	"fmt"
	"math/rand"
	// "strings"
)

// func LoadQuestion() map[string][]string {

// 	return map[string][]string{

// 		"Math": {

// 			"What is 2 + 2?\nA. 4\nB. 5",
// 			"\nA\n",

// 			"What is 10 - 7?\nA. 1\nB. 3",
// 			"\nB\n",

// 			"What is 5 x 6?\nA. 30\nB. 35",
// 			"\nA\n",

// 			"What is 12 / 3?\nA. 6\nB. 4",
// 			"\nB\n",
// 		},

// 		"English": {

// 			"What is the opposite of hot?\nA. Cold\nB. Warm",
// 			"\nA\n",

// 			"Choose the correct spelling.\nA. Recieve\nB. Receive",
// 			"\nB\n",

// 			"What is a noun?\nA. Action word\nB. Naming word",
// 			"\nB\n",

// 			"Choose the correct sentence.\nA. She go to school\nB. She goes to school",
// 			"\nB\n",
// 		},

// 		"Science": {

// 			"What planet do we live on?\nA. Mars\nB. Earth",
// 			"\nB\n",

// 			"What gas do humans breathe in?\nA. Oxygen\nB. Carbon Dioxide",
// 			"\nA\n",

// 			"What is H2O?\nA. Water\nB. Salt",
// 			"\nA\n",

// 			"How many legs does an insect have?\nA. 6\nB. 8",
// 			"\nA\n",
// 		},

// 		"Government": {

// 			"What arm of government makes laws?\nA. Legislature\nB. Judiciary",
// 			"\nA\n",

// 			"Who is the head of state in Nigeria?\nA. Governor\nB. President",
// 			"\nB\n",

// 			"What is democracy?\nA. Rule by the people\nB. Rule by soldiers",
// 			"\nA\n",

// 			"Nigeria practices what system?\nA. Monarchy\nB. Federalism",
// 			"\nB\n",
// 		},
// 	}
// }

// This distribute the number of questions equally among randomized selected subjects
// if there are remainders it assigns them to the early selected subject
func DistributeQuestions(subjects []string, total int) map[string]int {

	distribution := map[string]int{}

	base := total / len(subjects)

	remainder := total % len(subjects)
	newrem := 1
	// if remainder > 1 {
	// 	newrem = remainder / remainder
	// }
	// Shuffle the subjects
	rand.Shuffle(len(subjects), func(i, j int) {
		subjects[i], subjects[j] = subjects[j], subjects[i]
	})
	for i, subject := range subjects {
		if i < remainder {
			distribution[subject] = base + newrem
		} else {
			distribution[subject] = base
		}
	}

	return distribution
}

func RandomQuestion(subjects []string, numb int) []string {
	if len(subjects) == 0 {
		return nil
	}
	subjectMap := DistributeQuestions(subjects, numb)
	// This takes a slice of all the names of subjects that the user wants to answer, and also the number of questions that would be asked in the test.
	testQuestions := []string{}
	questionNbr := 1
	// It calls LoadQuestion() to get the entire questions in the data base.
	allQuestions := LoadQuestion()

	// Collect questions from selected subjects
	// Returns the slice of these questions and answers.
	for subject, nb := range subjectMap {
		questions, exists := allQuestions[subject]
		if !exists {
			continue
		}
		// Temporary slice copy for all selected subjects
		availableQuestions := make([]string, len(questions))
		copy(availableQuestions, questions)

		// Randomly pick questions
		// Questions should be at index 0,2,4,6 and so on.
		// While answers would be at odd indexes.
		testQuestions = append(testQuestions, subject, "\n")
		for i := 0; i < nb; i++ {
			// Stop if no more questions
			if len(availableQuestions) < 2 {
				break
			}
			// Generate random EVEN index
			randomIndex := rand.Intn(len(availableQuestions)/2) * 2

			// This give numbering to the question slices only
			// Append question and answer
			questionNumb := fmt.Sprintf("%d. ", questionNbr)
			testQuestions = append(
				testQuestions, questionNumb,
				availableQuestions[randomIndex],
				availableQuestions[randomIndex+1], "\n",
			)
			questionNbr++

			// Remove already selected question + answer
			availableQuestions = append(
				availableQuestions[:randomIndex],
				availableQuestions[randomIndex+2:]...,
			)
		}
	}
	return testQuestions
}

// func main() {
// 	subjects := []string{
// 		"Math",
// 		"Science",
// 		"English",
// 	}

// 	questions := RandomQuestion(subjects, 11)
// 	cleanquestions := strings.Join(questions, "")

// 	// fmt.Printf("%s\n", questions)
// 	fmt.Println(cleanquestions)
// }
