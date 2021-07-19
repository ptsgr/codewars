package reverseorrotate

import (
	"strconv"
)

func Revrot(s string, n int) string {
	res := ""
	if n <= 0 || s == "" || n > len(s) {
		return res
	}

	s = s[:len(s)-len(s)%n]
	for i := 0; i < len(s)/n; i++ {
		res += reverseChunk(s[i*n : (i+1)*n])
	}
	return res
}

func checkChunk(s string) bool {
	var res int
	for _, v := range s {
		tmp, _ := strconv.Atoi(string(v))
		res += tmp * tmp
	}
	if res%2 == 0 {
		return true
	}
	return false
}

func reverseStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseChunk(s string) string {
	if checkChunk(s) {
		return reverseStr(s)
	} else {
		return s[1:] + s[:1]
	}
}
