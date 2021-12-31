package leetcode

import (
	"testing"
)

type question26 struct {
	para26
	ans26
}

// para 是参数
// one 代表第一个参数
type para26 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans26 struct {
	one int
}

func Test_Problem26(t *testing.T) {

	qs := []question26{

		{
			para26{[]int{1, 1, 2}},
			ans26{2},
		},

		{
			para26{[]int{0, 0, 1, 1, 1, 1, 2, 3, 4, 4}},
			ans26{5},
		},

		{
			para26{[]int{0, 0, 0, 0, 0}},
			ans26{1},
		},

		{
			para26{[]int{1}},
			ans26{1},
		},
	}

	for _, q := range qs {
		w, p := q.ans26.one, q.para26
		a := removeDuplicates(p.one)
		if a != w {
			t.Errorf("【input】:%v\n【output】:%v\n【expect】:%v\n", p.one, a, w)
		}
	}
}
