package pkg

/*
	url: https://leetcode.com/explore/challenge/card/july-leetcoding-challenge-2021/608/week-1-july-1st-july-7th/3805/

		Given an n x n matrix where each of the rows and columns are sorted in ascending order, return the kth smallest element in the matrix.
	Note that it is the kth smallest element in the sorted order, not the kth distinct element.
*/

import "sort"

// KthSmallest find kth smallest
func KthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	arr := make([]int, n*n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			arr[i*n+j] = matrix[i][j]
		}
	}

	sort.Ints(arr)

	return arr[k-1]
}
