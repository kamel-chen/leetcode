package main

import (
	"fmt"
	"leetcode/array/pkg"
)

func main() {
	// a := [][]int{{1, 2}, {3, 4}}
	// fmt.Println(pkg.MatrixReshape(a, 4, 1))

	// b := []int{9, 77, 63, 22, 92, 9, 14, 54, 8, 38, 18, 19, 38, 68, 58, 19}
	// fmt.Println(pkg.MinSetSize(b))

	c := [][]int{{1, 2}, {1, 3}}
	fmt.Println(pkg.KthSmallest(c, 2))

}
