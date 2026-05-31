package main

import "project/cbt"

func main() {

	questions := cbt.LoadQuestion()

	cbt.Upload("Maths", questions)
}
