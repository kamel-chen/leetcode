package pkg

import (
	"math/rand"
)

/*
	url: https://leetcode.com/explore/challenge/card/july-leetcoding-challenge-2021/610/week-3-july-15th-july-21st/3820/

		Given an integer array nums, design an algorithm to randomly shuffle the array.
	All permutations of the array should be equally likely as a result of the shuffling.

	Implement the Solution class:
		Solution(int[] nums) Initializes the object with the integer array nums.
		int[] reset() Resets the array to its original configuration and returns it.
		int[] shuffle() Returns a random shuffling of the array.

*/

// Solution solution
type Solution struct {
	nums []int
}

// Constructor constructor
func Constructor(nums []int) Solution {
	return Solution{
		nums: nums,
	}
}

// Reset to origin arrays
func (s *Solution) Reset() []int {
	return s.nums
}

// Shuffle array
func (s *Solution) Shuffle() []int {
	a := make([]int, len(s.nums))
	copy(a, s.nums)

	// rand.Shuffle(len(a), func(i, j int) {
	// 	a[i], a[j] = a[j], a[i]
	// })

	for i := 0; i < len(a); i++ {
		j := rand.Intn(len(a)-i) + i
		a[i], a[j] = a[j], a[i]
	}

	return a
}
