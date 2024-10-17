package main

import "fmt"

func main() {

	var myNewInt int
	var myNewIntPointer *int
	myNewIntPointer = &myNewInt
	fmt.Println(myNewIntPointer)

	myInt := 8
	myIntPointer := &myInt
	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer)

	myNewestInt := 5
	fmt.Println(myNewestInt)
	myNewestIntPointer := &myNewestInt
	*myNewestIntPointer = 8
	fmt.Println(myNewestIntPointer)
	fmt.Println(myNewestInt)

}
