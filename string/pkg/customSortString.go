package pkg

/*
	url: https://leetcode.com/explore/challenge/card/july-leetcoding-challenge-2021/609/week-2-july-8th-july-14th/3813/

		order and str are strings composed of lowercase letters. In order, no letter occurs more than once.
	order was sorted in some custom order previously. We want to permute the characters of str so that they match the order that order was sorted.
	More specifically, if x occurs before y in order, then x should occur before y in the returned string.
	Return any permutation of str (as a string) that satisfies this property.

*/

// CustomSortString custom sort string
func CustomSortString(order string, str string) string {
	m := make([]int, 27)
	o := make([]string, 27)
	output := ""

	for i, v := range order {
		m[v-96] = i + 1
	}

	for _, v := range str {
		s := m[v-96]
		if s != 0 {
			o[s] += string(v)
		} else {
			o[0] += string(v)
		}
	}

	for i := 1; i < len(o); i++ {
		output += o[i]
	}

	output += o[0]

	return output
}
