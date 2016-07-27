// Source : https://leetcode.com/problems/two-sum/
// Author : Code.CC
// Date   : 2016-07-15

/**********************************************************************************
*
* Given an array of integers, find two numbers such that they add up to a specific target number.
*
* You may assume that each input would have exactly one solution.
*
* Input: numbers={2, 7, 11, 15}, target=9
* Output: index1=0, index2=1
*
*
**********************************************************************************/

package main

import (
	"fmt"
)

var (
	numbers []int
	target  int
	sum     int
)

func twoSum(numbers []int, target int) {

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			sum = numbers[i] + numbers[j]
			if sum == target {
				fmt.Printf("index1 = %d, index2 = %d", i, j)
				fmt.Println()
			}
		}
	}
}

func main() {
	numbers = []int{2, 7, 11, 15}
	target = 9

	twoSum(numbers, target)

}
