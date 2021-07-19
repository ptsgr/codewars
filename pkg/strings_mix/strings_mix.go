package stringsmix

import (
	"sort"
)

type result struct {
	str        string
	countRune  int
	compareStr int
}

type resultArray []result

func (a resultArray) Len() int { return len(a) }
func (a resultArray) Less(i, j int) bool {
	if a[i].countRune > a[j].countRune {
		return true

	} else if a[i].countRune == a[j].countRune {
		if a[i].compareStr < a[j].compareStr {
			return true
		} else if a[i].compareStr == a[j].compareStr {
			if a[i].str < a[j].str {
				return true
			}
		}
	}
	return false
}
func (a resultArray) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a resultArray) String() string {
	var resultStr string
	for _, v := range a {
		resultStr += v.str
	}
	if len(a) > 0 {
		return resultStr[:len(resultStr)-1]
	}
	return ""
}

func Mix(s1, s2 string) string {
	resM := make(map[rune]int)
	var res resultArray
	m1 := accountStr(s1)
	m2 := accountStr(s2)
	for k, v := range m1 {
		resM[k] += v
	}
	for k, v := range m2 {
		resM[k] += v
	}

	for k := range resM {
		if m1[k] > m2[k] {
			res.appendResult(1, m1[k], k)
		} else if m1[k] < m2[k] {
			res.appendResult(2, m2[k], k)
		} else {
			res.appendResult(3, m2[k], k)
		}
	}
	sort.Sort(res)

	return res.String()
}

func (r *resultArray) appendResult(stringNumber, runeCount int, symbol rune) {

	if runeCount < 2 {
		return
	}
	var resultStr string

	switch stringNumber {
	case 1:
		resultStr += "1:"
	case 2:
		resultStr += "2:"
	case 3:
		resultStr += "=:"
	}

	for i := 0; i < runeCount; i++ {
		resultStr += string(symbol)
	}
	resultStr += "/"
	*r = append(*r, result{resultStr, runeCount, stringNumber})

}

func accountStr(s string) (res map[rune]int) {
	res = make(map[rune]int)
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			res[c]++
		}
	}
	return
}
