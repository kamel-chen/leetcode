package pkg

/*
	url: https://leetcode.com/explore/challenge/card/august-leetcoding-challenge-2021/613/week-1-august-1st-august-7th/3836/
*/

// TwoSum find two sum
func TwoSum(nums []int, target int) []int {
	o := make([]int, 0, 2)
	m := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		r := target - nums[i]

		if m[r] != 0 && m[r] != i {
			o = append(o, i, m[r])
			break
		}
	}

	return o
}
