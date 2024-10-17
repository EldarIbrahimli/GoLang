package main

import (
	"fmt"
)

func main() {

	var amount, total float64

	fmt.Println("Welcome to the liter calculator")

	amount, err := paintNeeded(-5.2, 12)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("You need a total of %.2f liters for this wall\n\n", amount)
	}
	total += amount

	amount, err = paintNeeded(10.8, 19)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("You need a total of %.2f liters for this wall\n\n", amount)
	}
	total += amount

	/*amount, err = paintNeeded(11, 15.6)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("You need a total of %.2f liters for this wall\n\n", amount)
	}
	total += amount

	amount, err = paintNeeded(10, 23.3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("You need a total of %.2f liters for this wall\n\n", amount)
	}
	total += amount*/

	fmt.Printf("The total of liters needed is: %.2f liters\n", total)

}

func paintNeeded(width, height float64) (float64, error) {

	if width < 0 {
		return 0, fmt.Errorf("\nA width of %.2f is invalid\n", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("\nA height of %.2f is invalid\n", height)
	}

	area := width * height
	return area / 10, nil

}

/*func main() {

	var num1 int
	var num2 int

	fmt.Println("Input the first number")
	fmt.Scan(&num1)
	fmt.Println("Input the second number")
	fmt.Scan(&num2)

	sum1, multiply1, substract1 := sumMultiplyAndSubstract(num1, num2)

	fmt.Println("The summations is:\n", sum1, "\nThe multiplication is:\n", multiply1, "\nThe substraction is:\n", substract1)

}

/*func sum(a int, b int) int {

	sum := a + b
	return sum
}*/

/*func sumMultiplyAndSubstract(a int, b int) (int, int, int) {

	sum := a + b
	multiply := a * b
	substract := a - b
	return sum, multiply, substract
}*/
