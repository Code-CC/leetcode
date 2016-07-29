// Source : https://leetcode.com/problems/reverse-integer/
// Author : Code.CC
// Date   : 2016-07-28

/**********************************************************************************
*
* Determine whether an integer is a palindrome. Do this without extra space.
*
**********************************************************************************/

package main

import (
	"fmt"
)

func isPalindrome(numbers int) bool {
	var reminder int
	newNum := 0
	for i := numbers; i != 0; i = i / 10 {
		reminder = i % 10
		newNum = newNum*10 + reminder
	}
	if numbers == newNum {
		return true
	} else {
		return false
	}

}

func main() {
	numbers := -12321
	fmt.Printf("The numbers %d is a palindrome? ", numbers)
	fmt.Println(isPalindrome(numbers))
}
