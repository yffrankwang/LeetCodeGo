package leetcode

import (
	"testing"
)

type question20 struct {
	para20
	ans20
}

// para 是参数
// one 代表第一个参数
type para20 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans20 struct {
	one bool
}

func Test_Problem20(t *testing.T) {

	qs := []question20{

		{
			para20{"(([]){})"},
			ans20{true},
		},
		{
			para20{"()[]{}"},
			ans20{true},
		},
		{
			para20{"(]"},
			ans20{false},
		},
		{
			para20{"({[]})"},
			ans20{true},
		},
		{
			para20{"(){[({[]})]}"},
			ans20{true},
		},
		{
			para20{"((([[[{{{"},
			ans20{false},
		},
		{
			para20{"(())]]"},
			ans20{false},
		},
		{
			para20{""},
			ans20{true},
		},
		{
			para20{"["},
			ans20{false},
		},
	}

	for _, q := range qs {
		w, p := q.ans20, q.para20
		a := isValid(p.one)
		if w.one != a {
			t.Errorf("【input】:%v\n【output】:%v\n【expect】:%v\n", p, a, w.one)
		}
	}
}
