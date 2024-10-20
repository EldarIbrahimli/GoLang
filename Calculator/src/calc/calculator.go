package calculator

import "fmt"

func Calculate(num1 float64, operator string, num2 float64) (float64, error) {
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
			return 0, fmt.Errorf("error: division by zero")
		}
		result = num1 / num2
	default:
		return 0, fmt.Errorf("invalid operator. Supported operators are +, -, *, /")
	}
	return result, nil
}
