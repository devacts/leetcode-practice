package main

import "fmt"

func findlongestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}

	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		// Case 1: Odd length (center is a single char, e.g., "aba")
		len1 := expandAroundCenter(s, i, i)
		// Case 2: Even length (center is between two chars, e.g., "bb")
		len2 := expandAroundCenter(s, i, i+1)

		// Find the longer of the two expansions
		maxLen := len1
		if len2 > len1 {
			maxLen = len2
		}

		// If this is the longest we've seen, update our pointers
		if maxLen > end-start {
			// Find boundaries from center and length
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	// Return the length of the palindrome found
	// Formula: (right - 1) - (left + 1) + 1 => right - left - 1
	return right - left - 1
}

func main() {
	// s := "babad"
	// s := "xbaddaby"
	s := "xdddddy"
	// s := "dd"
	// s := "a"
	fmt.Printf("string is: %s\n", s)
	sub := findlongestPalindrome(s)
	fmt.Println(sub)
}
