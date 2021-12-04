package leetcode

import (
	"fmt"
	"testing"
)

type question8 struct {
	para8
	ans8
}

// para 是参数
// one 代表第一个参数
type para8 struct {
	one string
}

// ans 是答案
// one 代表第一个答案
type ans8 struct {
	one int
}

func Test_Problem8(t *testing.T) {

	qs := []question8{

		{
			para8{""},
			ans8{0},
		},

		{
			para8{"42"},
			ans8{42},
		},

		{
			para8{"   -42"},
			ans8{-42},
		},

		{
			para8{"4193 with words"},
			ans8{4193},
		},

		{
			para8{"words and 987"},
			ans8{0},
		},

		{
			para8{"-91283472332"},
			ans8{-2147483648},
		},

		{
			para8{"2147483646"},
			ans8{2147483646},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 8------------------------\n")

	for _, q := range qs {
		a, p := q.ans8, q.para8
		a1 := myAtoi(p.one)
		r := "OK"
		if a.one != a1 {
			r = "NG"
		}
		fmt.Printf("%s【input】:%v    【output】:%v\n", r, p.one, a1)
	}
	fmt.Printf("\n\n\n")
}
