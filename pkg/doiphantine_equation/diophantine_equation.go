package doiphantineequation

import (
	"math"
)

func Solequa(n int) [][]int {
	result := [][]int{}
	z := float64(n)
	for a := float64(1); a <= math.Sqrt(z); a++ {
		if math.Mod(z, a) == 0 {
			b := z / a
			if math.Mod((a+b), 2) == 0 && math.Mod((b-a), 4) == 0 {
				result = append(result, []int{int((a + b) / 2), int((b - a) / 4)})
			}
		}
	}

	return result
}
