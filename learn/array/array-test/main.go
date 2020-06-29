package main

import (
	"fmt"
	"leetcode/learn/array"
)

func main() {
	nums := []int{1, 1, 0, 1, 1, 1}
	max := array.FindMaxConsecutiveOnes(nums)
	fmt.Println(max)
}
