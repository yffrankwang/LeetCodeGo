package leetcode

func isMatch(s string, p string) bool {
	if p == "" {
		return s == p
	}

	if p == ".*" {
		return true
	}

	return deepMatch(p, s)
}

func deepMatch(p, s string) bool {
	for len(p) > 0 {
		pc := p[0]
		pn := byte(0)
		if len(p) > 1 {
			pn = p[1]
		}

		if pn == '*' {
			if pc == '.' {
				if deepMatch(p[2:], s) {
					return true
				}

				if len(s) > 0 {
					return deepMatch(p, s[1:])
				}
				return false
			}

			// skip pc
			s2 := s
			for s2 != "" {
				if pc == s2[0] {
					s2 = s2[1:]
					if deepMatch(p, s2) {
						return true
					}
				}
				break
			}

			p = p[2:]
			continue
		}

		if pc == '.' {
			if len(s) == 0 {
				return false
			}
			s = s[1:]
			p = p[1:]
			continue
		}

		if len(s) == 0 {
			return false
		}
		sc := s[0]
		if sc != pc {
			return false
		}
		s = s[1:]
		p = p[1:]
	}
	return len(s) == 0 && len(p) == 0
}
