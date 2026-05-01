package main

import "fmt"

func longestPalindrome(s string) string {
	var curr_longest_len int
	var curr_longest_pal, substr string
	curr_longest_pal = string(s[0])
	for i := 0; i <= len(s)-2; i = i + 1 {
		for j := i + 1; j <= i+2 && j < len(s); j++ {
			if s[i] == s[j] {
				substr = findLargerPallindrome(i-1, j+1, s)
				if len(substr) > curr_longest_len {
					curr_longest_len = len(substr)
					curr_longest_pal = substr
				}
			}
		}
	}
	return curr_longest_pal
}

func findLargerPallindrome(m, n int, str string) string {
	for ; m >= 0 && n < len(str); m, n = m-1, n+1 {
		if str[m] != str[n] {
			return str[m+1 : n]
		}
	}
	return str[m+1 : n]
}

func main() {
	// s := "babad"
	// s := "xbaddaby"
	s := "xdddddy"
	// s := "dd"
	// s := "a"
	fmt.Printf("string is: %s\n", s)
	sub := longestPalindrome(s)
	fmt.Println(sub)
}
