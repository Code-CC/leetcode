// Source : https://leetcode.com/problems/longest-substring-without-repeating-characters/
// Author : Code.CC
// Date   : 2016-07-20

/**********************************************************************************
*
* Given a string, find the length of the longest substring without repeating characters.
*
* Examples:
*
* Given "abcabcbb", the answer is "abc", which the length is 3.
*
* Given "bbbbb", the answer is "b", with the length of 1.
*
* Given "pwwkew", the answer is "wke", with the length of 3.
* Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*
*
**********************************************************************************/

package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	longest := 1
	var length int
	for i := 0; i < len(s); i++ {
		substring := make([]string, 0)
		substring = append(substring, string(s[i]))
		for j := i + 1; j < len(s); j++ {
			unique := true
			for k := 0; k < len(substring); k++ {
				if substring[k] == string(s[j]) {
					unique = false
					break
				}
			}
			if unique {
				substring = append(substring, string(s[j]))
				length = len(substring)
			} else {
				break
			}
		}

		if length > longest {
			longest = length
		}
	}

	return longest
}

func main() {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}
