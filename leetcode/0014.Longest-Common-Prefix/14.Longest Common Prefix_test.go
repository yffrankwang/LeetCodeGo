package leetcode

import (
	"testing"
)

type question14 struct {
	para14
	ans14
}

// para 是参数
type para14 struct {
	strs []string
}

// ans 是答案
type ans14 struct {
	ans string
}

func Test_Problem14(t *testing.T) {

	qs := []question14{

		{
			para14{[]string{"flower", "flow", "flight"}},
			ans14{"fl"},
		},

		{
			para14{[]string{"dog", "racecar", "car"}},
			ans14{""},
		},
	}

	for _, q := range qs {
		w, p := q.ans14, q.para14
		a := longestCommonPrefix(p.strs)
		if w.ans != a {
			t.Errorf("【input】:%v    【output】:%v\n", p.strs, a)
		}
	}
}
