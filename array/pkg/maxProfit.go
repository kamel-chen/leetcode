package pkg

import "math"

/**
 * https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/564/
 */

func BetterMaxProfit(prices []int) int {
	o := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			o += prices[i] - prices[i-1]
		}
	}

	return o
}

func MaxProfit(prices []int) int {
	var v int = math.MaxInt32
	var o int

	for _, p := range prices {		
		if p >= v {
			o += p - v
		}

		v = p
	}

	return o
}
