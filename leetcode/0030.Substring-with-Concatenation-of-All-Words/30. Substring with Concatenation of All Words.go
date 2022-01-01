package leetcode

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}

	res := []int{}

	sl := len(s)
	wl := len(words[0])
	wn := wl * len(words)

	for i := 0; i <= sl-wn; i++ {
		s1 := s[i : i+wn]
		if isMatchString(s1, wl, words) {
			res = append(res, i)
		}
	}

	return res
}

func isMatchString(s string, wl int, words []string) bool {
	wmap := map[string]int{}
	for _, w := range words {
		wmap[w]++
	}

	for i := 0; i < len(s); i += wl {
		s1 := s[i : i+wl]
		if _, ok := wmap[s1]; !ok {
			return false
		}
		wmap[s1]--
	}

	for _, v := range wmap {
		if v > 0 {
			return false
		}
	}

	return true
}
