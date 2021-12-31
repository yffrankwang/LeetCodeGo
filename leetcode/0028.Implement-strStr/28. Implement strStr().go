package leetcode

func strStr(haystack string, needle string) int {
	hl := len(haystack)
	nl := len(needle)

	if hl < nl {
		return -1
	}

	for i := 0; i < hl-nl+1; i++ {
		s := haystack[i : i+nl]
		if s == needle {
			return i
		}
	}

	return -1
}
