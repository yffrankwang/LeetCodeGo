package leetcode

import "sort"

func longestCommonPrefix(strs []string) string {
	return longestCommonPrefix1(strs)
}

func longestCommonPrefix1(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	var a []byte
	for i := 0; ; i++ {
		c := byte(0)
		for _, s := range strs {
			if i >= len(s) {
				return string(a)
			}

			if c == 0 {
				c = s[i]
				continue
			}

			if c != s[i] {
				return string(a)
			}
		}
		a = append(a, c)
	}
}

func longestCommonPrefix2(strs []string) string {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) <= len(strs[j])
	})
	minLen := len(strs[0])
	if minLen == 0 {
		return ""
	}
	var commonPrefix []byte
	for i := 0; i < minLen; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != strs[0][i] {
				return string(commonPrefix)
			}
		}
		commonPrefix = append(commonPrefix, strs[0][i])
	}
	return string(commonPrefix)
}
