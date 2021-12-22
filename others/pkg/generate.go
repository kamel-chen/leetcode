package pkg

/**
 * https://leetcode.com/explore/interview/card/top-interview-questions-easy/99/others/601/
 */

func Generate(numRows int) [][]int {
	o := [][]int{}
	o = append(o, []int{1})

	for r := 2; r <= numRows; r++ {
		tmp := []int{}
		b := r-2
		for i := 0; i < r; i++ {
			if i == 0 {
				tmp = append(tmp, o[b][0])
				continue
			}
			if i == r-1 {
				tmp = append(tmp, o[b][b])
				continue
			}
			tmp = append(tmp, o[b][i] + o[b][i-1])
		}

		o = append(o, tmp)
	}

    return o
}
