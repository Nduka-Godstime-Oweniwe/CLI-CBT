package cbt

import (
	"encoding/json"
	"os"
)

type Score struct {
	ID      int
	Subject string
	Score   string
}

type Question struct {
	ID       int
	Subject  string
	Question string
	OptionA  string
	OptionB  string
	Answer   string
}

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
func LoadQuestion() map[string][]Question {
	// This function goes to a file named "questions.json" and takes the data inside an returns it
	// if the data does not exist return an empty data type map[string][]Question
	result := map[string][]Question{}
	data, err := os.ReadFile("questions.json")
	if err != nil {
		data, _ = json.Marshal(result)
		err = os.WriteFile("questions.json", data, 0644)
	} else {
		err = json.Unmarshal(data, &result)
	}
	return result
}
func LoadScore() map[string]Score {
	// This function goes to a file named "scores.json" and takes the data inside an returns it
	// if the data does not exist return an empty data type map[string]Score
	result := map[string]Score{}
	data, err := os.ReadFile("scores.json")
	if err != nil {
		data, _ = json.Marshal(result)
		err = os.WriteFile("scores.json", data, 0644)
	} else {
		err = json.Unmarshal(data, &result)
	}
	return result
}

func DumpQuestion(questions map[string][]Question) {
	// This function takes in the questions as parameters and overwrites what was inside questions.json with the new questions
	// If the file doesnt exist , the function creates it and dumps it inside
	result := map[string][]Question{}
	data, err := json.Marshal(questions)
	if err == nil {
		os.WriteFile("questions.json", data, 0644)
	} else {
		data, _ = json.Marshal(result)
		os.WriteFile("questions.json", data, 0644)
	}

}

func DumpScore(scores map[string]Score) {
	// This function takes in the scores as parameters and overwrites what was inside scores.json with the new scores
	// If the file doesnt exist , the function creates it and dumps it inside
	result := map[string]Score{}
	data, err := json.Marshal(scores)
	if err == nil {
		os.WriteFile("scores.json", data, 0644)
	} else {
		data, _ = json.Marshal(result)
		os.WriteFile("scores.json", data, 0644)
	}
}
