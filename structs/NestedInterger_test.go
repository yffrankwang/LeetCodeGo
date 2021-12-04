package structs

import (
	"testing"
)

func Test_NestedInteger(t *testing.T) {
	n := NestedInteger{}

	assertTrue(t, n.IsInteger())

	n.SetInteger(1)
	assertEqual(t, 1, n.GetInteger())

	elem := NestedInteger{Num: 1}

	expected := NestedInteger{
		Num: 1,
		Ns:  []*NestedInteger{&elem},
	}
	n.Add(elem)

	assertEqual(t, expected, n)

	assertEqual(t, expected.Ns, n.GetList())
}
