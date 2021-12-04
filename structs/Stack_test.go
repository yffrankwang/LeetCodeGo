package structs

import (
	"testing"
)

func Test_Stack(t *testing.T) {
	s := NewStack()
	assertTrue(t, s.IsEmpty(), "检查新建的 s 是否为空")

	start, end := 0, 100

	for i := start; i < end; i++ {
		s.Push(i)
		assertEqual(t, i-start+1, s.Len(), "Push 后检查 q 的长度。")
	}

	for i := end - 1; i >= start; i-- {
		assertEqual(t, i, s.Pop(), "从 s 中 pop 出数来。")
	}

	assertTrue(t, s.IsEmpty(), "检查 Pop 完毕后的 s 是否为空")
}
