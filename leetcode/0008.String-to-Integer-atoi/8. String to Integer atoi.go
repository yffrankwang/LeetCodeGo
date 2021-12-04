package leetcode

func myAtoi(s string) int {
	i := 0
	z := len(s)

	// skip space
	for ; i < z; i++ {
		if s[i] != ' ' {
			break
		}
	}

	if i >= z {
		return 0
	}

	// find sign
	var sign int64 = 1
	switch s[i] {
	case '+':
		i++
	case '-':
		sign = -1
		i++
	}

	// skip zeor
	for ; i < z; i++ {
		if s[i] != '0' {
			break
		}
	}

	// find last
	j := i
	for ; j < z; j++ {
		c := s[j]
		if c < '0' || c > '9' {
			break
		}
	}

	if j <= i {
		return 0
	}

	// max
	var max int64
	if sign > 0 {
		max = int64(2<<30) - 1
	} else {
		max = int64(2 << 30)
	}

	// convert
	var n int64
	var p int64 = 1
	for k := j - 1; k >= i; k-- {
		n += int64(s[k]-'0') * p
		p *= 10
		if n > max || p > 10000000000 {
			n = max
			break
		}
	}

	return int(int64(sign) * n)
}
