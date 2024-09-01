package main

import (
	"fmt"
)

func searchAndPrint(target int) {
	// Create the sorted array from 1 to 32
	arr := make([]int, 32)

	// Fill the array using a traditional for loop
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}

	// Binary search
	low, high := 0, len(arr)-1
	var index int
	found := false

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == target {
			index = mid
			found = true
			break
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if found {
		fmt.Printf("Number %d found at index %d.\n", target, index)
	} else {
		fmt.Println("Number not found in the array.")
	}
}

func main() {
	target := 27
	searchAndPrint(target)
}
