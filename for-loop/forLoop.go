package main

import "fmt"

func main() {

	/* for i := 0; i <= 5; i++ { // create an i integer variable, if the variable is less than or equal to  5, add 1

		fmt.Println(i) // prints the variable

	} */

	/* for i := 0; i <= 100; i += 10 { // the same loop but now increases by ten till its 100

		fmt.Println(i) // prints the variable

	} */

	/* for i := 0; i <= 5; i++ { // create an i integer variable, if the variable is less than or equal to  5, add 1

			if i == 3 { // if the variable becomes 3
				continue // skip it
			}
			if i == 4 { // if the variable is 4
	            break // break/end the code
	        }

			fmt.Println(i) // print the variable

		} */

	/* how := [2]string{"big", "tasty"} // creates an array

	fruits := [3]string{"apple", "orange", "banana"} // creates a second array

	for i := 0; i < len(how); i++ { // this whole code creates an inner loop

		for j := 0; j < len(fruits); j++ {

			fmt.Println(how[i], fruits[j]) // prints them until both for loops are satisfied
		}
	} */

	//fruits := [3]string{"apple", "orange", "banana"} // creates an array

	//for idx, val := range fruits { // creates a range for loop to make using arrays easier
	//	fmt.Printf("%v\t%v\n", idx, val) // prints the index and value of an array
	//}

	// for _, val := range fruits { // prints without index, can be also used for value
	// 	fmt.Printf("%v\n", val) // prints just the value of an array
	// }

	/*z := 0
	n := 2

	for i := 1; i <= n; i++ {

		fmt.Println("I is equal to", i)

		for j := 1; j <= n; j++ {

			fmt.Println("\t j is equal to", j)

			for k := 1; k <= n; k++ {

				fmt.Println("\t\t k is equal to", k)

				z++

			}
		}
	}

	fmt.Println("z is equal to", z)
	*/
	/*for x := 1; x <= 3; x++{
	  	fmt.Println(x)
	  }
	  for (x := 1; x <= 3; x++{ // adding parantheless after for is not allowed
	  	fmt.Println(x))
	  }
	  for x := 1; x = 3; x++{ // will only go up to 2
	  	fmt.Println(x)
	  }
	  for x := 1; x <= 3; x++{ // not in the scope of for so will just print the initial number of x
	  }
	  fmt.Println(x)
	*/
	for x := 1; x >= 3; x++ {
		fmt.Println(x)
	}
}
