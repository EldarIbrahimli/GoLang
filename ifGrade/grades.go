package main

import "fmt"

func main() {
	var AliMathGrade int
	var AliAzGrade int
	var AliSportGrade int

	fmt.Println("Type the math grade:")
	fmt.Scan(&AliMathGrade)
	fmt.Println("Type the Azerbaijan grade:")
	fmt.Scan(&AliAzGrade)
	fmt.Println("Type the Sports grade:")
	fmt.Scan(&AliSportGrade)

	if AliMathGrade <= 2 {

		fmt.Println("Ali failed math")

	} else {
		fmt.Println("Ali passed Math")

		if AliAzGrade <= 2 {

			fmt.Println("Ali failed Azerbaijan")

		} else {

			fmt.Println("Ali passed Azerbaijani Language")

			if AliSportGrade <= 2 {

				fmt.Println("Ali failed Sports")

			} else {

				fmt.Println("Ali passed the Sports")

			}

		}
	}

}
