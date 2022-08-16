package main

import "fmt"

// Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.
// Example 1:

// Input: s = "leetcode"
// Output: 0
// Example 2:

// Input: s = "loveleetcode"
// Output: 2
// Example 3:

// Input: s = "aabb"
// Output: -1

func firstUniqChar(s string) int {
	d := map[byte]int{}

	// Count each character.
	for i := 0; i < len(s); i++ {
		d[s[i]]++
	}

	// Find the first unique character and return.
	for i := 0; i < len(s); i++ {
		if d[s[i]] == 1 {
			return i
		}
	}

	// If there's no unique character then return -1.
	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
}
