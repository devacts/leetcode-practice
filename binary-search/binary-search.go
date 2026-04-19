package main

import "fmt"

func binarySearch(s []int, target int) int {
	length := len(s)
	left := 0
	right := length - 1

	for mid := (length - 1) / 2; right >= left; mid = left + (right-left)/2 {
		switch {
		case target > s[mid]:
			left = mid + 1
		case target < s[mid]:
			right = mid - 1
		default:
			return mid
		}
	}
	return -1
}

func main() {
	numSlice := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	// target := 23
	// target := 91
	// target := 2
	target := 8

	pos := binarySearch(numSlice, target)
	fmt.Printf("Position of target %d in slice is %d\n", target, pos)
}
