package leetcode

import (
	"reflect"
	"testing"

	"github.com/yffrankwang/LeetCodeGo/structs"
)

type question25 struct {
	para25
	ans25
}

// para 是参数
// one 代表第一个参数
type para25 struct {
	one []int
	two int
}

// ans 是答案
// one 代表第一个答案
type ans25 struct {
	one []int
}

func Test_Problem25(t *testing.T) {

	qs := []question25{

		{
			para25{
				[]int{1, 2, 3, 4, 5},
				2,
			},
			ans25{[]int{2, 1, 4, 3, 5}},
		},

		{
			para25{
				[]int{1, 2, 3, 4, 5},
				3,
			},
			ans25{[]int{3, 2, 1, 4, 5}},
		},

		{
			para25{
				[]int{1, 2, 3, 4, 5},
				1,
			},
			ans25{[]int{1, 2, 3, 4, 5}},
		},
	}

	for _, q := range qs {
		w, p := q.ans25.one, q.para25
		a := structs.List2Ints(reverseKGroup(structs.Ints2List(p.one), p.two))
		if !reflect.DeepEqual(w, a) {
			t.Errorf("【input】:%v\n【output】:%v\n【expect】:%v\n", p, a, w)
		}
	}
}
