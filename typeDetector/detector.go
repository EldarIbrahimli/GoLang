package main

import "fmt"

func main() {

	var whatewer string

	fmt.Println("Write a number, string or a float")
	fmt.Scan(&whatewer)

	fmt.Printf("your variable is an %T", whatewer)

}
