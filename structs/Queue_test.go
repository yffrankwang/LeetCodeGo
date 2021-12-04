package structs

import (
	"testing"
)

func Test_Queue(t *testing.T) {
	q := NewQueue()
	assertTrue(t, q.IsEmpty(), "检查新建的 q 是否为空")

	start, end := 0, 100

	for i := start; i < end; i++ {
		q.Push(i)
		assertEqual(t, i-start+1, q.Len(), "Push 后检查 q 的长度。")
	}

	for i := start; i < end; i++ {
		assertEqual(t, i, q.Pop(), "从 q 中 pop 出数来。")
	}

	assertTrue(t, q.IsEmpty(), "检查 Pop 完毕后的 q 是否为空")
}
