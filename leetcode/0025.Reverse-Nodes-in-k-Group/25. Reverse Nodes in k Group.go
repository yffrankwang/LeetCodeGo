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
func reverseKGroup(head *ListNode, k int) *ListNode {
	root := &ListNode{Next: head}

	prev := root
	for {
		next := prev.Next
		for i := 0; i < k; i++ {
			if next == nil {
				return root.Next
			}
			next = next.Next
		}

		for node, i := prev.Next, 0; i < k; i++ {
			node, node.Next, next = node.Next, next, node
		}
		prev, prev.Next = prev.Next, next
	}
}
