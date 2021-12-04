package leetcode

import (
	"fmt"
	"testing"
)

type question6 struct {
	para6
	ans6
}

// para 是参数
// one 代表第一个参数
type para6 struct {
	s       string
	numRows int
}

// ans 是答案
// one 代表第一个答案
type ans6 struct {
	one string
}

func Test_Problem6(t *testing.T) {

	qs := []question6{
		{
			para6{"PAYPALISHIRING", 1},
			ans6{"PAYPALISHIRING"},
		},

		{
			para6{"PAYPALISHIRING", 2},
			ans6{"PYAIHRNAPLSIIG"},
		},

		{
			para6{"PAYPALISHIRING", 3},
			ans6{"PAHNAPLSIIGYIR"},
		},

		{
			para6{"PAYPALISHIRING", 4},
			ans6{"PINALSIGYAHRPI"},
		},

		{
			para6{"A", 1},
			ans6{"A"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 6------------------------\n")

	for _, q := range qs {
		a, p := q.ans6, q.para6
		as := convert(p.s, p.numRows)
		r := "OK"
		if a.one != as {
			r = "NG"
		}
		fmt.Printf("%s【input】:%v       【output】:%v\n", r, p, as)
	}
	fmt.Printf("\n\n\n")
}
