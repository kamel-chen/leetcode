package pkg

/**
https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/549/
*/

func singleNumber(nums []int) int {
	m := map[int]bool{}
	for _, num := range nums {
		if m[num] {
			delete(m, num)
		} else {
			m[num] = true
		}
	}

	v := 0
	for k := range m {
		v = k
	}

	return v
}

func betterSingleNumber(nums []int) int {
	r := 0
	for _, n := range nums {
		r = r ^ n
	}

	return r
}
