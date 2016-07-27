// Source : https://leetcode.com/problems/string-to-integer-atoi/
// Author : Code.CC
// Date   : 2016-07-27

/**********************************************************************************
*
* Implement atoi to convert a string to an integer.
*
**********************************************************************************/

package main

import (
	"fmt"
)

func myAtoi(s string) int {
	numbers := 0
	if s[0] == '-' {
		for i := 1; i < len(s); i++ {
			if int(s[i]) - 48 {

			}
			numbers = numbers*10 - (int(s[i]) - 48)
		}
	} else {
		for i := 0; i < len(s); i++ {
			numbers = numbers*10 + (int(s[i]) - 48)
		}
	}

	return numbers
}

func main() {
	s := "5623"
	x := "-295"
	fmt.Println(myAtoi(s))
	fmt.Println(myAtoi(x))
}
