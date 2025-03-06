package main

import "math"

// ReverseArray reverses the elements of an array in place
func ReverseArray(arr *[5]int) {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}

// MaxMinInSlice returns the maximum and minimum values in a slice
func MaxMinInSlice(s []int) (int, int) {
	if len(s) == 0 {
		return 0, 0
	}

	max := math.MinInt
	min := math.MaxInt

	for _, v := range s {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max, min
}

// RemoveAtIndex returns a new slice with the element at the specified index removed
func RemoveAtIndex(s []int, index int) []int {
	if index < 0 || index >= len(s) {
		return s
	}

	result := make([]int, 0, len(s)-1)
	result = append(result, s[:index]...)
	result = append(result, s[index+1:]...)
	return result
}

// RemoveElement returns a new slice with all occurrences of the specified element removed
func RemoveElement(s []int, element int) []int {
	result := make([]int, 0, len(s))

	for _, v := range s {
		if v != element {
			result = append(result, v)
		}
	}

	return result
}

// RemoveFirstElement returns a new slice with the first occurrence of the specified element removed
func RemoveFirstElement(s []int, element int) []int {
	result := make([]int, 0, len(s))
	found := false

	for _, v := range s {
		if !found && v == element {
			found = true
			continue
		}
		result = append(result, v)
	}

	return result
}
