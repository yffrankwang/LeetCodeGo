package leetcode

// 3 rows
// P   A   H   N   -> step: 4
// A P L S I I G   -> step: 2, 2
// Y   I   R       -> step: 4
// -----------------------------
// 4 rows
// P     I    N    -> step: 6
// A   L S  I G    -> step: 4, 2
// Y A   H R       -> step: 2, 4
// P     I         -> step: 6
func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	z := len(s)

	rs := make([]byte, 0, z)

	smax := (numRows - 1) << 1

	var steps [2]int

	step := 0
	for i := 0; i < numRows; i++ {
		steps[0] = (numRows - i - 1) << 1
		steps[1] = smax - steps[0]
		if steps[0] == 0 {
			steps[0] = smax
		}
		if steps[1] == 0 {
			steps[1] = smax
		}

		for j, k := i, 0; j < z; j += step {
			rs = append(rs, s[j])
			step = steps[k&1]
			k++
		}
	}
	return string(rs)
}

// solution 2: follow zigzag to save character in matrix
func convert2(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	bs := make([][]byte, numRows, numRows) // matix
	d := true                              // is down or up
	r := 0                                 // row

	for i := 0; i < len(s); i++ {
		bs[r] = append(bs[r], s[i])
		if d {
			if r == numRows-1 {
				r--
				d = false
			} else {
				r++
			}
		} else {
			if r == 0 {
				r++
				d = true
			} else {
				r--
			}
		}
	}

	rs := make([]byte, 0, len(s))
	for _, r := range bs {
		for _, b := range r {
			rs = append(rs, b)
		}
	}
	return string(rs)
}
