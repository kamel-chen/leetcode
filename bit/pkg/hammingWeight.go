package pkg

import "math/bits"

/**
 * https://leetcode.com/explore/interview/card/top-interview-questions-easy/99/others/565/
 */

func HammingWeight(num uint32) int {
    o := 0

	for num > 0 {
		if num % 2 == 1 {
			o ++
		}
		num = num >> 1
	}
	return o
}

func BuiltInFunc(num uint32) int {
	return bits.OnesCount32(num)
}
