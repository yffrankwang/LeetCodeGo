package leetcode

import (
	"reflect"
	"testing"
)

type question18 struct {
	para18
	ans18
}

// para 是参数
// one 代表第一个参数
type para18 struct {
	a []int
	t int
}

// ans 是答案
// one 代表第一个答案
type ans18 struct {
	one [][]int
}

func Test_Problem18(t *testing.T) {

	qs := []question18{

		{
			para18{[]int{-1, 0, 1, 2, -1, -4}, -1},
			ans18{[][]int{{-4, 0, 1, 2}, {-1, -1, 0, 1}}},
		},

		{
			para18{[]int{1, 1, 1, 1}, 4},
			ans18{[][]int{{1, 1, 1, 1}}},
		},

		{
			para18{[]int{0, 1, 5, 0, 1, 5, 5, -4}, 11},
			ans18{[][]int{{-4, 5, 5, 5}, {0, 1, 5, 5}}},
		},

		{
			para18{[]int{1, 0, -1, 0, -2, 2}, 0},
			ans18{[][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		},

		{
			para18{[]int{1, 0, -1, 0, -2, 2, 0, 0, 0, 0}, 0},
			ans18{[][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}, {0, 0, 0, 0}}},
		},

		{
			para18{[]int{1, 0, -1, 0, -2, 2, 0, 0, 0, 0}, 1},
			ans18{[][]int{{-2, 0, 1, 2}, {-1, 0, 0, 2}, {0, 0, 0, 1}}},
		},
	}

	for i, q := range qs {
		w, p := q.ans18, q.para18
		a := fourSum(p.a, p.t)
		if !reflect.DeepEqual(w.one, a) {
			t.Errorf("%d【input】:%v\n【output】:%v\n【expect】:%v\n", i, p, a, w.one)
		}
	}
}
