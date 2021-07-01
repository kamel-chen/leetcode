package binarytree

/*
	url: https://leetcode.com/explore/challenge/card/june-leetcoding-challenge-2021/607/week-5-june-29th-june-30th/3797/

	content:
		Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.
	According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q
	as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”
*/

// TreeNode definition for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isParent(p, q *TreeNode) bool {
	if p.Left == nil && p.Right == nil {
		return false
	}

	if p.Left == q || p.Right == q {
		return true
	}

	if p.Left == nil {
		return isParent(p.Right, q)
	} else if p.Right == nil {
		return isParent(p.Left, q)
	} else {
		return isParent(p.Left, q) || isParent(p.Right, q)
	}
}

// LowestCommonAncestor find LCA
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if isParent(p, q) {
		return p
	}

	if isParent(q, p) {
		return q
	}

	f := true

	for f {
		if root.Left != nil && isParent(root.Left, p) && isParent(root.Left, q) {
			root = root.Left
			continue
		} else if root.Right != nil && isParent(root.Right, p) && isParent(root.Right, q) {
			root = root.Right
			continue
		}

		break
	}

	return root
}
