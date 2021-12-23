package pkg

import "strconv"

/**
 * https://leetcode.com/explore/interview/card/top-interview-questions-easy/102/math/743/
 */

func FizzBuzz(n int) []string {
	var o []string

	for i := 1; i <= n; i++ {
		if (i % 3 == 0) && (i % 5 == 0) {
			o = append(o, "FizzBuzz")
			continue	
		}
		if i % 3 == 0 {
			o = append(o, "Fizz")
			continue
		}
		if i % 5 == 0 {
			o = append(o, "Buzz")
			continue
		}
		o = append(o, strconv.Itoa(i))
	}

	return o
}
