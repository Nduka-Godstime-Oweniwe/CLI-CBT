package cbt

import (
	"math/rand"
)

// This distribute the number of questions equally among randomized selected subjects
// if there are remainders it assigns them to the early selected subject
func DistributeQuestions(subjects []string, total int) map[string]int {

	if len(subjects) == 0 {
		return nil
	}

	distribution := map[string]int{}

	base := total / len(subjects)

	remainder := total % len(subjects)
	newrem := 1
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

func RandomQuestion(subjects []string, numb int) []Question {
	if len(subjects) == 0 {
		return nil
	}
	subjectMap := DistributeQuestions(subjects, numb)
	// This takes a slice of all the names of subjects that the user wants to answer, and also the number of questions that would be asked in the test.
	testQuestions := []Question{}

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
		availableQuestions := make([]Question, len(questions))
		copy(availableQuestions, questions)

		// Randomly pick questions from the available questions for the subject
		rand.Shuffle(
			len(availableQuestions),
			func(i, j int) {
				availableQuestions[i],
					availableQuestions[j] =
					availableQuestions[j],
					availableQuestions[i]
			},
		)
		if nb > len(availableQuestions) {
			nb = len(availableQuestions)
		}
		testQuestions = append(
			testQuestions,
			availableQuestions[:nb]...,
		)
	}
	return testQuestions
}
