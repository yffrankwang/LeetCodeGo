package leetcode

func reverse(x int) int {
	var n int64 = 0
	for x != 0 {
		n = n*10 + int64(x%10)
		x = x / 10
	}
	if n > 1<<31-1 || n < -(1<<31) {
		return 0
	}
	return int(n)
}
