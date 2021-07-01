package main

import (
	"fmt"
)

/**
	link: https://leetcode.com/explore/challenge/card/june-leetcoding-challenge-2021/607/week-5-june-29th-june-30th/3796/

	content:
		Given a binary array nums and an integer k, return the maximum number of consecutive 1's in the array if you can flip at most k 0's.

	Example 1:

	Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
	Output: 6
	Explanation: [1,1,1,0,0,1,1,1,1,1,1]
	Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
*/

func main() {
	a := []int{0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1}

	fmt.Println(longestOnes(a, 3))
}

func longestOnes(nums []int, k int) int {
	max := 0
	l := 0
	window := []int{}

	getMax := func() {
		if (l > max) {
			max = l
		}
	}

    for _, num := range nums {
		if (num == 1) {
			l++
			window = append(window, num)
			getMax()
			continue
		}

		// num = 0
		if (k != 0) {
			k--
			l++
			window = append(window, num)
			getMax()
			continue
		}

		// num = 0 & k = 0
		for i := 0; i < len(window); i++ {
			if (window[i] == 0) {
				window = window[i+1:]
				window = append(window, num)
				break
			} else {
				l--
				if (i == (len(window) - 1)) {
					window = []int{}
				}
			}
		}

		getMax()
	}

	return max
}
