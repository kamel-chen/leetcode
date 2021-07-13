package pkg

/*
	url: https://leetcode.com/explore/challenge/card/july-leetcoding-challenge-2021/609/week-2-july-8th-july-14th/3811/

		Given two strings s and t, determine if they are isomorphic.
	Two strings s and t are isomorphic if the characters in s can be replaced to get t.
	All occurrences of a character must be replaced with another character while preserving the order of characters.
	No two characters may map to the same character, but a character may map to itself.
*/

// IsIsomorphic is isomorphic
func IsIsomorphic(s string, t string) bool {
	memo := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		if memo[s[i]] == 0 {
			for _, v := range memo {
				if v == t[i] {
					return false
				}
			}

			memo[s[i]] = t[i]
			continue
		}

		if memo[s[i]] != t[i] {
			return false
		}
	}

	return true
}
