package main

import "fmt"

func main() {
	amount := 6
	amountAdress := &amount
	fmt.Println(&amountAdress)
	fmt.Println(amountAdress)
	double(&amount)

	fmt.Println(amount)
}
func double(number *int) {
	*number += 2

}
