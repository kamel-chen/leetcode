package pkg

/**
 * https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/627/
 */

// Definition for a binary tree node.
// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func IsSymmetric(root *TreeNode) bool {
	var recur func(l *TreeNode, r *TreeNode) bool

	recur = func(l *TreeNode, r *TreeNode) bool {
		if (l == nil && r == nil) {
			return true
		}
		if ((l == nil || r == nil) || (l.Val != r.Val)) {
			return false
		}

		return recur(l.Left, r.Right) && recur(l.Right, r.Left)
	}

	return recur(root.Left, root.Right)
}
