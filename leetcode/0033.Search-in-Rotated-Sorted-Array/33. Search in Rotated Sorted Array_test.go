package leetcode

import (
	"testing"
)

type question33 struct {
	para33
	ans33
}

// para 是参数
// one 代表第一个参数
type para33 struct {
	nums   []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans33 struct {
	one int
}

func Test_Problem33(t *testing.T) {

	qs := []question33{

		{
			para33{[]int{3, 1}, 1},
			ans33{1},
		},

		{
			para33{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
			ans33{4},
		},

		{
			para33{[]int{4, 5, 6, 7, 0, 1, 2}, 3},
			ans33{-1},
		},
	}

	for _, q := range qs {
		w, p := q.ans33.one, q.para33
		a := search(p.nums, p.target)
		if w != a {
			t.Errorf("【input】:%v\n【output】:%v\n【expect】:%v\n", p, a, w)
		}
	}
}
