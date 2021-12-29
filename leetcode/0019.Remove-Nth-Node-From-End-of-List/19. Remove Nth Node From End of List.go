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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	root := &ListNode{Next: head}

	slow, fast := root, head
	for fast != nil {
		if n > 0 {
			n--
		} else {
			slow = slow.Next
		}
		fast = fast.Next
	}

	if n == 0 {
		slow.Next = slow.Next.Next
	}
	return root.Next
}
