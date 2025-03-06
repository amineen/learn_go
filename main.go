package main

import (
	"fmt"
	"math"
)

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	fmt.Println("Before:", numbers)
	// reverseArray(&numbers)
	// fmt.Println("After:", numbers)

	max, min := maxMinInSlice(numbers[:])
	fmt.Println("Max:", max)
	fmt.Println("Min:", min)

	numbers2 := []int{1, 2, 3, 4, 5, 6}

	numbers3 := removeAtIndex(numbers2, 2)
	fmt.Println("Number 2:", numbers2)
	fmt.Println("Number 3:", numbers3)
}

func reverseArray(arr *[5]int) {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

func maxMinInSlice(s []int) (int, int) {
	max := math.MinInt
	min := math.MaxInt

	for i := range s {
		if s[i] >= max {
			max = s[i]
		}
		if s[i] <= min {
			min = s[i]
		}
	}
	return max, min
}
func removeAtIndex(s []int, index int) []int {
	if index < 0 || index >= len(s) {
		return s
	}

	result := make([]int, 0, len(s)-1)
	result = append(result, s[:index]...)
	result = append(result, s[index+1:]...)
	return result
}

func removeElement(s []int, element int) []int {

	result := make([]int, 0, len(s))

	for _, v := range s {
		if v != element {
			result = append(result, v)
		}
	}

	return result
}
