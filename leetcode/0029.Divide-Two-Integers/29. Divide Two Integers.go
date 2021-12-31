package leetcode

func divide(dividend int, divisor int) int {
	minus := false
	if dividend < 0 || divisor < 0 {
		if dividend < 0 && divisor < 0 {
			dividend = -dividend
			divisor = -divisor
		} else {
			if dividend < 0 {
				dividend = -dividend
			}
			if divisor < 0 {
				divisor = -divisor
			}
			minus = true
		}
	}
	if dividend < divisor {
		return 0
	}

	result := 0
	if divisor == 1 {
		result = dividend
	} else {
		divide := 0
		index := getBits(dividend)
		for index > 0 {
			val := getBit(dividend, index)
			index--
			divide = (divide << 1) + val

			if divide < divisor {
				result <<= 1
			} else {
				result = (result << 1) + 1
				divide = divide - divisor
			}
		}
	}

	if minus {
		return -result
	}

	if result > 2147483647 {
		result = 2147483647
	}
	return result
}

func getBits(num int) int {
	numLen := 0
	for num != 0 {
		numLen++
		num = num >> 1
	}
	return numLen
}

func getBit(num, pos int) int { //获取从右到左的第pos位置的值1/0
	mask := 1 << (pos - 1)

	if (num & mask) == mask {
		return 1
	}
	return 0
}
