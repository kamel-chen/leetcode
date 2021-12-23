package pkg

/**
 * https://leetcode.com/explore/interview/card/top-interview-questions-easy/97/dynamic-programming/569/
 */

func RecurClimbStairs(n int) int {
	if (n == 1) {
		return 1
	}
	if (n == 2) {
		return 2
	}

	return RecurClimbStairs(n-1) + RecurClimbStairs(n-2)
}

func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	}

	o := make([]int, n)
	o[0] = 1
	o[1] = 2

	for i := 2; i < n; i++ {
		o[i] = o[i-1] + o[i-2]
	}

	return o[n-1]
}

func BetterClimbStairs(n int) int {
	if n <= 2 {
		return n
	}

	sum := 0
	n1 := 1
	n2 := 2

	for i := 2; i < n; i++ {
		sum = n1 + n2
		n1 = n2
		n2 = sum
	}

	return sum
}
