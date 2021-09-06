package weightforweight

import (
	"sort"
	"strconv"
	"strings"
)

func OrderWeight(strng string) (result string) {
	weights := make(map[int][]string)
	var keys []int
	split := strings.Split(strng, " ")
	for _, w := range split {
		nums := strings.Split(w, "")
		var sum int
		for _, n := range nums {
			res, err := strconv.Atoi(n)
			if err != nil {
				continue
			}
			sum += res
		}
		weights[sum] = append(weights[sum], w)
	}
	for k := range weights {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		sort.Strings(weights[k])
		for _, res := range weights[k] {
			result += res + " "
		}

	}
	result = result[:len(result)-1]
	return
}
