package leetcode

import (
	"testing"
)

type question29 struct {
	para29
	ans29
}

// para 是参数
// one 代表第一个参数
type para29 struct {
	dividend int
	divisor  int
}

// ans 是答案
// one 代表第一个答案
type ans29 struct {
	one int
}

func Test_Problem29(t *testing.T) {

	qs := []question29{

		{
			para29{-2147483648, -1},
			ans29{2147483647},
		},

		{
			para29{10, 3},
			ans29{3},
		},

		{
			para29{7, -3},
			ans29{-2},
		},

		{
			para29{-1, 1},
			ans29{-1},
		},

		{
			para29{1, -1},
			ans29{-1},
		},

		{
			para29{2147483647, 3},
			ans29{715827882},
		},
	}

	for _, q := range qs {
		w, p := q.ans29.one, q.para29
		a := divide(p.dividend, p.divisor)
		if w != a {
			t.Errorf("【input】:%v\n【output】:%v\n【expect】:%v\n", p, a, w)
		}
	}
}
