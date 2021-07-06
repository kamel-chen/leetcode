package main

import (
	"fmt"
	"leetcode/array/pkg"
)

func main() {
	a := [][]int{{1, 2}, {3, 4}}

	fmt.Println(pkg.MatrixReshape(a, 4, 1))
}
