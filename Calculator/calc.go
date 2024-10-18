package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome To Simple Calculator")
	fmt.Println("Enter calculations like 5 + 10 or 10 / 2. Type '$' to exit.")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		// ASK: Why does $ exit not work with that
		// TODO: Ask chatgpt how to write these sentences in a corret way
		input = strings.TrimSpace(input)

		if input == "$" {
			fmt.Println("Exiting calculator.")
			break
		}

		var num1, num2 float64
		var operator string

		_, err = fmt.Sscanf(input, "%f %s %f", &num1, &operator, &num2)
		if err != nil {
			fmt.Println("Invalid input. Please enter in the format 'number operator number'.")
			continue
		}

		var result float64
		switch operator {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			if num2 == 0 {
				fmt.Println("Error: Division by zero.")
				continue
			}
			result = num1 / num2
		default:
			fmt.Println("Invalid operator. Supported operators are +, -, *, /.")
			continue
		}

		fmt.Printf("Result: %.2f\n", result) // the f does not let the program output scientific response
		//fmt.Printf("Result: %v\n", result)
	}
}
