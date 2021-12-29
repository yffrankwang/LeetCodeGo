package leetcode

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	opens := []byte{}
	for i := 0; i < len(s); i++ {
		v := s[i]
		if v == '[' || v == '(' || v == '{' {
			opens = append(opens, v)
		} else {
			if len(opens) == 0 {
				return false
			}

			o := opens[len(opens)-1]
			opens = opens[0 : len(opens)-1]
			if !((o == '[' && v == ']') || (o == '(' && v == ')') || (o == '{' && v == '}')) {
				return false
			}
		}
	}
	return len(opens) == 0
}
