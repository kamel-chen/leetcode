package pkg

/*
	url: https://leetcode.com/explore/challenge/card/july-leetcoding-challenge-2021/608/week-1-july-1st-july-7th/3803/

		In MATLAB, there is a handy function called reshape which can reshape an m x n matrix into a new one with a different size r x c keeping its original data.
	You are given an m x n matrix mat and two integers r and c representing the row number and column number of the wanted reshaped matrix.
	The reshaped matrix should be filled with all the elements of the original matrix in the same row-traversing order as they were.
	If the reshape operation with given parameters is possible and legal, output the new reshaped matrix; Otherwise, output the original matrix.

*/

// MatrixReshape reshape matrix
func MatrixReshape(mat [][]int, r int, c int) [][]int {
	arr := []int{}
	o := [][]int{}
	var key int

	m := len(mat)
	n := len(mat[0])

	if m*n != r*c {
		return mat
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			arr = append(arr, mat[i][j])
		}
	}

	for k := 0; k < r; k++ {
		key = k * c
		o = append(o, arr[key:key+c])
	}

	return o
}
