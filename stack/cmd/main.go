package main

import (
	"fmt"
	"leetcode/stack/pkg/stack"
)

/*
	Problem url: https://leetcode.com/explore/challenge/card/june-leetcoding-challenge-2021/606/week-4-june-22nd-june-28th/3794/

	Content:
		You are given a string s consisting of lowercase English letters.
	A duplicate removal consists of choosing two adjacent and equal letters and removing them.
	We repeatedly make duplicate removals on s until we no longer can.
	Return the final string after all such duplicate removals have been made.
	It can be proven that the answer is unique.
*/

func main() {
	fmt.Println(removeDuplicates("azxxzy"))
}

func removeDuplicates(s string) string {
	sk := stack.New()

	for _, char := range s {
		ch := string(char)

		if sk.IsEmpty() {
			sk.Push(ch)
			continue
		}

		if el := sk.Pop(); el != ch {
			sk.Push(el)
			sk.Push(ch)
		}
	}

	return sk.String()
}
