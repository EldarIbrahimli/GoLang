package main

import "fmt"

func main() {
	math := result(5, 3) // creates a result function with 5 and 3

	fmt.Println(math) // prints it

	words := myWords("Hello", " ", "World")
	fmt.Println(words)
}

func myWords(str1, str2, str3 string) string {
	return str1 + str2 + str3
}

func result(num1, num2 int) int { // function for result is created
	return num1 + num2 // should return 8
}
