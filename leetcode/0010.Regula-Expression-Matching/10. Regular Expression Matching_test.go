package leetcode

import (
	"fmt"
	"testing"
)

type question10 struct {
	para10
	ans10
}

// para 是参数
// one 代表第一个参数
type para10 struct {
	one, two string
}

// ans 是答案
// one 代表第一个答案
type ans10 struct {
	one bool
}

func Test_Problem10(t *testing.T) {

	qs := []question10{

		{
			para10{"aa", "a"},
			ans10{false},
		},

		{
			para10{"aa", "a*"},
			ans10{true},
		},

		{
			para10{"ab", ".*"},
			ans10{true},
		},

		{
			para10{"aab", "c*a*b"},
			ans10{true},
		},

		{
			para10{"mississippi", "mis*is*p*."},
			ans10{false},
		},

		{
			para10{"aaa", "a*a"},
			ans10{true},
		},

		{
			para10{"aaaaaaaaaaaaab", "a*a*a*a*a*a*a*a*a*a*c"},
			ans10{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 10------------------------\n")

	for _, q := range qs {
		a, p := q.ans10, q.para10
		a1 := isMatch(p.one, p.two)
		r := "OK"
		if a1 != a.one {
			r = "NG"
		}
		fmt.Printf("%s【input】:%v, %v    【output】:%v\n", r, p.one, p.two, a1)
	}
	fmt.Printf("\n\n\n")
}
