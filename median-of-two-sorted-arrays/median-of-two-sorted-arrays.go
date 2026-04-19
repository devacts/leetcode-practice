package main

import (
	"fmt"
	"slices"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = slices.Concat(nums1, nums2)
	slices.Sort(nums1)
	length := len(nums1)
	var median float64

	if length%2 == 0 {
		median = float64(nums1[(length/2)-1]+nums1[length/2]) / 2
	} else {
		median = float64(nums1[(length-1)/2])
	}
	return median
}

func main() {
	// nums1 := []int{1, 2, 3}
	// nums2 := []int{4, 5}

	// nums1, nums2 := []int{1, 3}, []int{2}

	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	median := findMedianSortedArrays(nums1, nums2)
	fmt.Println(median)
}
