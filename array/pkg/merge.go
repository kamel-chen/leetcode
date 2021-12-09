package pkg

/**
 * https://leetcode.com/explore/interview/card/top-interview-questions-easy/96/sorting-and-searching/587/
 */
func Merge(nums1 []int, m int, nums2 []int, n int)  {
	if n == 0 {
		return
	}
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}

	for i := m+n-1; i >= 0; i-- {
		if m == 0 {
			nums1[i] = nums2[n-1]
			n--
			continue
		}
		if n == 0 {
			nums1[i] = nums1[m-1]
			m--
			continue
		}
		if nums1[m-1] > nums2[n-1] {
			nums1[i] = nums1[m-1]
			m--
		} else {
			nums1[i] = nums2[n-1]
			n--
		}
	}
}
