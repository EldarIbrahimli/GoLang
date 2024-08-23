package main

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

}
