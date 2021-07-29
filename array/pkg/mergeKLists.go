package pkg

import "fmt"

/*
	url: https://leetcode.com/problems/merge-k-sorted-lists/

*/

type ListNode struct {
    Val int
    Next *ListNode
}

// PrintNode print all list node
func PrintNode(n *ListNode) {
	if n == nil {
		fmt.Println(n)
		return
	}
	for {
		fmt.Print(n.Val)
		if n.Next == nil {
			break
		}

		n = n.Next		
	}
	fmt.Print("\n")
}

// MergeKLists merge k lists
func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	m := []*ListNode{}

	for i:=0; i < len(lists); i+=2 {
		if i == len(lists) - 1 {
			m = append(m, lists[i])
			return MergeKLists(m)
		}

		node := &ListNode{}
		s := node
		l := lists[i]
		r := lists[i+1]

		if l == nil && r == nil {
			m = append(m, nil)
			continue
		}

		for {
			if l == nil {
				node.Val = r.Val
				r = r.Next
			} else if r == nil {
				node.Val = l.Val
				l = l.Next
			} else {
				if l.Val >= r.Val {
					node.Val = r.Val
					r = r.Next
				} else {
					node.Val = l.Val
					l = l.Next
				}
			}

			if l == nil && r == nil {
				break
			}
			node.Next = &ListNode{}
			node = node.Next
		}

		m = append(m, s)
	}

	return MergeKLists(m)
}
