package ipvalidation

const big = 0xFFFFFF

func dtoi(s string) (n int, i int, ok bool) {
	n = 0
	for i = 0; i < len(s) && '0' <= s[i] && s[i] <= '9'; i++ {
		if i == 0 && s[i] == '0' && len(s) > 2 && '0' <= s[i+1] && s[i+1] <= '9' {
			return 0, 0, false
		}
		n = n*10 + int(s[i]-'0')
		if n >= big {
			return big, i, false
		}
	}
	if i == 0 {
		return 0, 0, false
	}
	return n, i, true
}

func Is_valid_ip(s string) bool {
	for i := 0; i < 4; i++ {
		if len(s) == 0 {
			return false
		}
		if i > 0 {
			if s[0] != '.' {
				return false
			}
			s = s[1:]
		}
		n, c, ok := dtoi(s)
		if !ok || n > 0xFF {
			return false
		}
		s = s[c:]
	}
	if len(s) != 0 {
		return false
	}
	return true
}
