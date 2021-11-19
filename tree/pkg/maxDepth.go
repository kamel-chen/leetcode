package pkg

/**
 * https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/555/
 */

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func BetterMaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := BetterMaxDepth(root.Left) + 1
	r := BetterMaxDepth(root.Right) + 1

	if l > r {
		return l
	}
	return r
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := 1

	return dfs(root, max)
}

func dfs(t *TreeNode, m int) int {
	if t.Left != nil {
		l := dfs(t.Left, m + 1)
		if t.Right != nil {
			r := dfs(t.Right, m + 1)
			if r > l {
				return r
			} else {
				return l
			}
		}

		return l
	} else if t.Right != nil {
		return dfs(t.Right, m + 1)
	}

	return m
}
