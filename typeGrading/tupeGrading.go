package main

import (
	"fmt"
)

func main() {

	var grade int
	var status string

	fmt.Println("Welcome to the student grading")
	fmt.Println("Write the grade of a student:")

	// Read the input and check for errors

	_, err := fmt.Scan(&grade)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	if grade == 100 {
		status = "passed"
		fmt.Println("A perfect score of", grade, ",you have", status)
	} else if grade >= 60 {
		status = "passed"
		//fmt.Println("You passed the test with the grade of", grade)
	} else {
		status = "failed"
		//fmt.Println("You failed the test with the grade of", grade)
	}
	fmt.Println("You have a score of", grade, ", so you have", status, "the test")
}
