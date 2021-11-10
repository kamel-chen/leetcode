package pkg

/**
 * https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/882/
 */

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)

	for _, v := range s {
		m[v]++
	}

	for _, v := range t {
		m[v]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}
