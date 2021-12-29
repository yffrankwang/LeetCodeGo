package leetcode

var result []string
var keypad = [][]string{
	{},
	{},
	{"a", "b", "c"},
	{"d", "e", "f"},
	{"g", "h", "i"},
	{"j", "k", "l"},
	{"m", "n", "o"},
	{"p", "q", "r", "s"},
	{"t", "u", "v"},
	{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	result = []string{}
	if digits == "" {
		return result
	}

	letterCombinate("", digits)
	return result
}

func letterCombinate(res string, digits string) {
	if digits == "" {
		result = append(result, res)
		return
	}

	k := digits[0]
	digits = digits[1:]
	letters := keypad[k-'0']
	for i := 0; i < len(letters); i++ {
		res += letters[i]
		letterCombinate(res, digits)
		res = res[0 : len(res)-1]
	}
}
