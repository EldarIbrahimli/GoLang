package main

import "fmt"

func main() {

	for i := 2; i <= 6; i = i + 2 {
		fmt.Printf("%d\t", i+1)
	}

	//peppers("green", 13)  // call the function
	//peppers("yellow", 16) // call the function
	//peppers("red", 15)    // call the function

}

/*func peppers(color string, number int) { // created a variable and an integer inside a function

	fmt.Println("I need", number, "of", color, "peppers") // printed the functions

}*/

/*package main
import ("fmt")

func myFunction(x int, y int) (result int) {
  result = x + y
  return result
}

func main() {
  total := myFunction(1, 2)
  fmt.Println(total)
}*/

/*package main
import ("fmt")

func myFunction(x int, y int) int {
  return x + y
}

func main() {
  fmt.Println(myFunction(1, 2))
}*/
