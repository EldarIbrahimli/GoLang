package main

import (
	"fmt"
)

func linearSearch(arr []int, target int) int {
	for i, value := range arr {
		if value == target {
			return i // Return the index of the target
		}
	}
	return -1 // Return -1 if the target is not found
}

func main() {
	// Create a sorted array from 1 to 32
	arr := make([]int, 32)
	for i := 0; i < 32; i++ {
		arr[i] = i + 1
	}

	target := 27
	index := linearSearch(arr, target)
	if index != -1 {
		fmt.Printf("Element %d found at index %d\n", target, index)
	} else {
		fmt.Println("Element not found")
	}
}
