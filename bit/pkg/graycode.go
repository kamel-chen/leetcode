package pkg

/**
url : https://leetcode.com/explore/challenge/card/july-leetcoding-challenge-2021/608/week-1-july-1st-july-7th/3799/

content:

	An n-bit gray code sequence is a sequence of 2n integers where:

Every integer is in the inclusive range [0, 2n - 1],
The first integer is 0,
An integer appears no more than once in the sequence,
The binary representation of every pair of adjacent integers differs by exactly one bit, and
The binary representation of the first and last integers differs by exactly one bit.
Given an integer n, return any valid n-bit gray code sequence.

*/

// Graycode print sequence
func Graycode(n int) []int {
	o := []int{0, 1}
	l := 1

	for l != n {
		x := 1 << l
		l++

		for i := 0; i < x; i++ {
			o = append(o, o[x-1-i]+x)
		}

	}

	return o
}
