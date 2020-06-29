package array

/***

An Array is a collection of items.
The items could be integers, strings, DVDs, games, books—anything really.
The items are stored in neighboring (contiguous) memory locations.
Because they're stored together, checking through the entire collection of items is straightforward.

***/

/***

Given a binary array, find the maximum number of consecutive 1s in this array.

Example 1:
Input: [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s.
The maximum number of consecutive 1s is 3.

Note:

The input array will only contain 0 and 1.
The length of input array is a positive integer and will not exceed 10,000

***/
func FindMaxConsecutiveOnes(nums []int) int {
	var max, same int
	for i := 0; i < len(nums); i++ {

		if i == 0 {
			same = 1
			max = same
			continue
		}

		if nums[i] != nums[i-1] {
			if max <= same {
				max = same
			}
			same = 1
			continue
		}

		same++
	}

	if max <= same {
		max = same
	}

	return max
}
