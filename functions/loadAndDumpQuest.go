package cbt

// data :=map[string][]string{
// 	"Maths" : {
// 		"1+1 = ?\nA. 2 \nB. 0",
// 		"A",
// 		"1-1 =?\nA. 2\nB. 0",
// 		"B",
// 	},

// 	"Physics" : {
// 		"kg = ?\nA. mass \nB. force",
// 		"A",
// 		"Newton =?\nA. mass\nB. force",
// 		"B",
// 	}

// }

// The example above shows the format of how the questions and the answers would be stored.
// A map of key: string, and Value: slice of strings.
// The keys are the name of the subjects
// The values are the questions as (even indexes) and answers as odd indexes
func LoadQuestion() map[string][]string {
	// This function goes to a file named "questions.json" and takes the data inside an returns it
	// if the data does not exist return an empty data type map[string][]string

	return map[string][]string{}
}

func DumpQuestion(questions map[string][]string) {
	// This function takes in the questions as parameters and overwrites what was inside questions.json with the new questions
	// If the file doesnt exist , the function creates it and dumps it inside

}
