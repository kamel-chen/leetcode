package binarytree

/*
	url: https://leetcode.com/explore/challenge/card/july-leetcoding-challenge-2021/610/week-3-july-15th-july-21st/3819/

		Given a binary search tree (BST), find the lowest common ancestor (LCA) of two given nodes in the BST.
	According to the definition of LCA on Wikipedia:
		“The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”
*/

// TreeNode tree node
// type TreeNode struct {
//     Val   int
//     Left  *TreeNode
//     Right *TreeNode
// }

// LowestCommonAncestor2 find ancestor
func LowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	if (root.Val < p.Val && root.Val > q.Val) || (root.Val < q.Val && root.Val > p.Val) {
		return root
	}

	if root.Val > p.Val {
		root = root.Left
		return LowestCommonAncestor2(root, p, q)
	}

	root = root.Right
	return LowestCommonAncestor2(root, p, q)
}
