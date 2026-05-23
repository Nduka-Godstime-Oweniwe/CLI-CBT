package cbt

func Upload(subject string, questions map[string][]string) {

 // Keeps running until user stops
 for {

  // Gets question and answer from user
  question, answer := UserQuestion()

  // Adds question into the correct subject
  questions[subject] = append(questions[subject], question)
  questions[subject] = append(questions[subject], answer)

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

This was the first solution you gave me for upload.go 
