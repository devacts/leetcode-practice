package main

import "fmt"

func length_of_longest_substring(s string) int {
	charMap := make(map[rune]int)
	maxLen := 0
	left := 0
	for right, char := range s {
		lastIndex, ok := charMap[char]
		if ok && lastIndex >= left {
			left = lastIndex + 1
		}
		charMap[char] = right
		currentWindowLen := right - left + 1
		if currentWindowLen > maxLen {
			maxLen = currentWindowLen
		}
	}
	return maxLen
}

func main() {
	str := "abcabcbb"
	result := length_of_longest_substring(str)
	fmt.Printf("result is %d\n", result)
}
