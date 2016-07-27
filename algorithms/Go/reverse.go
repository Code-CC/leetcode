// Source : https://leetcode.com/problems/reverse-integer/
// Author : Code.CC
// Date   : 2016-07-26

/**********************************************************************************
*
* Reverse digits of an integer.
*
* Example1: x = 123, return 321
*
* Example2: x = -123, return -321
*
**********************************************************************************/

package main

import (
	"fmt"
)

func reverse(numbers int) int {
	var reminder int
	newNum := 0
	for i := numbers; i != 0; i = i / 10 {
		reminder = i % 10
		newNum = newNum*10 + reminder
	}
	return newNum
}

func main() {
	numbers := -12345
	fmt.Printf("The numbers is %d,The result is %d", numbers, reverse(numbers))
}
