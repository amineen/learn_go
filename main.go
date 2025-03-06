package main

import (
	"fmt"
)

func main() {
	// Example of array operations
	numbers := [5]int{2, 4, 6, 8, 10}
	fmt.Println("Before:", numbers)
	ReverseArray(&numbers)
	fmt.Println("After:", numbers)

	// Example of finding max and min values
	max, min := MaxMinInSlice(numbers[:])
	fmt.Println("Max:", max)
	fmt.Println("Min:", min)

	// Example of removing element at index
	numbers2 := []int{1, 2, 3, 4, 5, 6}
	numbers3 := RemoveAtIndex(numbers2, 2)
	fmt.Println("Original:", numbers2)
	fmt.Println("After removing index 2:", numbers3)

	// Example of removing all occurrences of an element
	numbers4 := []int{1, 2, 3, 2, 4, 5, 2}
	numbers5 := RemoveElement(numbers4, 2)
	fmt.Println("Original:", numbers4)
	fmt.Println("After removing all 2s:", numbers5)

	// Example of removing first occurrence of an element
	numbers6 := RemoveFirstElement(numbers4, 2)
	fmt.Println("After removing first 2:", numbers6)
}
