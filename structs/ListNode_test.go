package structs

import (
	"testing"
)

func Test_l2s(t *testing.T) {
	assertEqual(t, []int{}, List2Ints(nil), "输入nil，没有返回[]int{}")

	one2three := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	assertEqual(t, []int{1, 2, 3}, List2Ints(one2three), "没有成功地转换成[]int")

	limit := 100
	overLimitList := Ints2List(make([]int, limit+1))
	assertPanics(t, func() { List2Ints(overLimitList) }, "转换深度超过 %d 限制的链条，没有 panic", limit)
}

func Test_s2l(t *testing.T) {
	assertNil(t, Ints2List([]int{}), "输入[]int{}，没有返回nil")

	ln := Ints2List([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	i := 1
	for ln != nil {
		assertEqual(t, i, ln.Val, "对应的值不对")
		ln = ln.Next
		i++
	}
}

func Test_getNodeWith(t *testing.T) {
	//
	ln := Ints2List([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	val := 10
	node := &ListNode{
		Val: val,
	}
	tail := ln
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = node
	expected := node
	actual := ln.GetNodeWith(val)
	assertEqual(t, expected, actual)
}

func Test_Ints2ListWithCycle(t *testing.T) {
	ints := []int{1, 2, 3}
	l := Ints2ListWithCycle(ints, -1)
	assertEqual(t, ints, List2Ints(l))

	l = Ints2ListWithCycle(ints, 1)
	assertPanics(t, func() { List2Ints(l) })
}
