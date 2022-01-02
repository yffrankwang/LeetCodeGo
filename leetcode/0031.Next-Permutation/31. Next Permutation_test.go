package leetcode

import (
	"reflect"
	"testing"
)

type question31 struct {
	para31
	ans31
}

// para 是参数
// one 代表第一个参数
type para31 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans31 struct {
	one []int
}

func Test_Problem31(t *testing.T) {

	qs := []question31{

		{
			para31{[]int{1, 3, 2}},
			ans31{[]int{2, 1, 3}},
		},

		{
			para31{[]int{1, 2, 3}},
			ans31{[]int{1, 3, 2}},
		},

		{
			para31{[]int{3, 2, 1}},
			ans31{[]int{1, 2, 3}},
		},

		{
			para31{[]int{1, 1, 5}},
			ans31{[]int{1, 5, 1}},
		},

		{
			para31{[]int{1}},
			ans31{[]int{1}},
		},
	}

	for _, q := range qs {
		w, p := q.ans31.one, q.para31

		ns := make([]int, len(p.nums))
		copy(ns, p.nums)
		nextPermutation(p.nums)

		if !reflect.DeepEqual(w, p.nums) {
			t.Errorf("【input】:%v\n【output】:%v\n【expect】:%v\n", ns, p.nums, w)
		}
	}
}
