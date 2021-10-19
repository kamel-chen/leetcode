package pkg

/**
 * url: https://leetcode.com/explore/featured/card/top-interview-questions-easy/92/array/727/
 */

func RemoveDuplicates(nums []int) int {
	var i, v int

	if len(nums) == 0 {
		return 0
	}

	for _, v = range nums {
		if nums[i] != v {
			nums[i+1] = v
			i ++
		}
	}

	return i+1
}

func WorseRemoveDuplicates(nums []int) int {
	if (len(nums) == 0) {
		return 0
	}

    o := len(nums)
	num := nums[0]
	cur := 1

	for i := 1; i < len(nums); i++ {
		// fmt.Println(o, num, cur, i)
		if num == nums[i] {
			o --
		} else {
			num = nums[i]
			nums[cur] = num
			cur++
		}
	}

	return o
}
