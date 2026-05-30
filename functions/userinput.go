package cbt

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// //This function takes a string as Parameter.
// // E.g "Enter Question: " , "Select An Option"
// func UserInput(prompt string) string {
// 	return ""
// }

// // This function takes two parameters prompt and limit. The prompt would be passed into Input()
// //Input would be set into a loop until user types an int that is between 1 and limit
// // If user doesnt type an int between 1 and limit print "Invalid Option"
// // convert the input to int and return it
// func UserOption(prompt string, limit int) int {
// 	return 0
// }

func UserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)
	return userInput

}

func UserOption(prompt string, limit int) int {
	for {
		str, err := strconv.Atoi(UserInput(prompt))
		if err != nil || str == 0 || str > limit {
			fmt.Println("Invalid Option")
			continue
		}

		if str <= limit {
			return str
		}
	}

}
