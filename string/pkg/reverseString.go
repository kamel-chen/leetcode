package pkg

/**
 * https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/879/
*/

func reverseString(s []byte)  {
	var b byte

	for i:=0; i<len(s)/2; i++ {
		b = s[i]
		s[i] = s[len(s)-i-1]
		s[len(s)-i-1] = b
	}
}
