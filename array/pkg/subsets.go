package pkg

/*
	url: https://leetcode.com/problems/subsets/
*/

// Subsets find all subsets
func Subsets(nums []int) [][]int {
	o := [][]int{[]int{}}

	for i := 0; i < len(nums); i++ {
		l := len(o)
		for j := 0; j < l; j++ {

			tmp := append(o[j], nums[i])
			o = append(o, tmp)
		}
	}

	return o
}
