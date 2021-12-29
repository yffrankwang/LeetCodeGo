package leetcode

func isMatch(s string, p string) bool {
	return isMatch2(s, p)
}

// Approach 1: Recursion
func isMatch1(s string, p string) bool {
	if p == "" {
		return s == p
	}

	firstMatch := (s != "" && (p[0] == s[0] || p[0] == '.'))

	if len(p) > 1 && p[1] == '*' {
		if isMatch1(s, p[2:]) {
			return true
		}

		if firstMatch && isMatch1(s[1:], p) {
			return true
		}

		return false
	}

	return firstMatch && isMatch1(s[1:], p[1:])
}

// Approach 2: Dynamic Programming
func isMatch2(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[len(s)][len(p)] = true

	for i := len(s); i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			firstMatch := (i < len(s) && (p[j] == s[i] || p[j] == '.'))

			if j+1 < len(p) && p[j+1] == '*' {
				dp[i][j] = dp[i][j+2] || firstMatch && dp[i+1][j]
			} else {
				dp[i][j] = firstMatch && dp[i+1][j+1]
			}
		}
	}

	return dp[0][0]
}
