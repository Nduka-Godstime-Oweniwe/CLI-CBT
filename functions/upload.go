package cbt

func Upload(subject string, questions map[string][]Question) {

	// Keeps running until user stops
	for {

		// Gets question and answer from user
		question, optionA, optionB, answer := UserQuestion()

		// Adds question into the correct subject
		questions[subject] = append(questions[subject], Question{
			Question: question,
			OptionA:  optionA,
			OptionB:  optionB,
			Answer:   answer,
		})

		// Ask user if they want to continue
		option := UserOption("1. Add another question\n2. Go back\n", 2)

		// Stop loop if user chooses 2
		if option == 2 {
			break
		}
	}

	// Save updated questions into questions.json
	DumpQuestion(questions)
}
