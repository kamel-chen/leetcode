package pkg

/*
	url: https://leetcode.com/explore/challenge/card/july-leetcoding-challenge-2021/608/week-1-july-1st-july-7th/3804/

		Given an array arr.  You can choose a set of integers and remove all the occurrences of these integers in the array.
	Return the minimum size of the set so that at least half of the integers of the array are removed.
*/

// Set a set
type Set struct {
	mem   int
	count int
}

// MinSetSize find min set size
func MinSetSize(arr []int) int {
	set := make(map[int]int)
	l := len(arr)
	output := 0
	sum := 0

	for i := 0; i < l; i++ {
		set[arr[i]]++
	}

	for k, v := range set {
		if v == 1 {
			delete(set, k)
		}
	}

	for {
		if len(set) == 0 {
			output += (l/2 - sum)
			break
		}

		output++
		max := 0
		key := 0

		for k, v := range set {
			if v > max {
				max = v
				key = k
			}
		}

		delete(set, key)
		sum += max

		if sum >= l/2 {
			break
		}

	}

	return output
}
