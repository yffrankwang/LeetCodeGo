package leetcode

import "github.com/yffrankwang/LeetCodeGo/structs"

// ListNode define
type ListNode = structs.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	last := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			last.Next = l1
			last = l1
			l1 = l1.Next
		} else {
			last.Next = l2
			last = l2
			l2 = l2.Next
		}
	}

	last.Next = nil
	if l1 != nil {
		last.Next = l1
	}
	if l2 != nil {
		last.Next = l2
	}

	return head.Next
}
