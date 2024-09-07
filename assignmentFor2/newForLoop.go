package main

import "fmt"

func main() {

	fmt.Println("Sums of items up to 5 are:")

	var sum int

	for i := 1; i <= 5; i++ {

		sum += i

		count := 5 - i

		fmt.Println(sum)

		if count > 0 {

			fmt.Println("there are", count, "more numbers")

		} else {

			fmt.Println("that is all!")

		}
	}

}
