package pkg

import "sort"

/**
 * https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/578/
 */

func ContainsDuplicate(nums []int) bool {
	if (len(nums) < 2) {
		return false
	}

	sort.Ints(nums)
	var tmp = nums[0];

	for i:=1; i<len(nums); i++ {
		if (nums[i] == tmp) {
		return true
		}
		tmp = nums[i]
	}

	return false
}

func BetterContainsDuplicate(nums []int) bool {
	m := make(map[int]bool)

	for _, n := range nums {
	  if m[n] {
		return true
	  }
  
	  m[n] = true
	}
  
	return false
}
