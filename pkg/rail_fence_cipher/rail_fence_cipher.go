package railfencecipher

type index struct {
	up   bool
	rail int
}

func printRails(r [][]rune) (res string) {
	for _, str := range r {
		res += string(str)
	}
	return
}

func splitStr(s string, n int) [][]rune {
	res := make([][]rune, n)
	i := index{true, 0}

	for _, r := range s {

		res[i.rail] = append(res[i.rail], r)

		if i.rail == n-1 {
			i.up = false
		}
		if i.rail == 0 {
			i.up = true
		}

		if i.up {
			i.rail++
		} else {
			i.rail--
		}
	}
	return res
}

func Encode(s string, n int) string {
	return printRails(splitStr(s, n))
}

func Decode(s string, n int) string {
	rails := splitStr(s, n)
	slen := len(s)
	res := []rune{}
	for index, rail := range rails {
		rails[index] = []rune(s[:len(rail)])
		s = s[len(rail):]
	}
	i := index{true, 0}

	for x := 0; x < slen; x++ {
		res = append(res, rails[i.rail][0])
		rails[i.rail] = rails[i.rail][1:]

		if i.rail == n-1 {
			i.up = false
		}
		if i.rail == 0 {
			i.up = true
		}

		if i.up {
			i.rail++
		} else {
			i.rail--
		}

	}
	return string(res)
}
