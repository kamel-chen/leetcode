package main

import (
	"fmt"
	"leetcode/string/pkg"
)

func main() {
	s := "badc"
	t := "baba"
	fmt.Println(pkg.IsIsomorphic(s, t))
}
