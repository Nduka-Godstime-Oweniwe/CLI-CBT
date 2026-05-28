package cbt

// import "fmt"

func Upload(subject string, questions map[string][]string) {
	// fmt.Println(questions)
	for {
		question, answer := UserQuestion()
		questions[subject] = append(questions[subject], question)
		questions[subject] = append(questions[subject], answer)
		// fmt.Println(questions)
		option := UserOption("1. Add Another question\n2. Go back\n", 2)
		if option == 2 {
			DumpQuestion(questions)
			break
		}
	}
	//  for {
	//   question, answer := UserQuestion()
	//   questions[subject] = append(questions[subject], question)
	//   questions[subject] = append(questions[subject], answer)

	//   option := UserOption("1. Add another question\n2. Go back\n", 2)

	//   if option == 2 {
	// 	DumpQuestion(questions)
	//    	break
	//   }
	//  }
	//
}
