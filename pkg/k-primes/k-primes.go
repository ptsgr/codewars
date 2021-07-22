package kprimes

func primeFind(n int) (result int) {
	i := 2
	for i*i <= n {
		for n%i == 0 {
			n /= i
			result++
		}
		i++
	}
	if n > 1 {
		result++
	}
	return
}

func CountKprimes(k, start, nd int) []int {
	var NPrimeSlice []int

	for i := start; i <= nd; i++ {
		if primeFind(i) == k {
			NPrimeSlice = append(NPrimeSlice, i)
		}
	}
	return NPrimeSlice
}

func Puzzle(s int) (result int) {
	for _, c := range CountKprimes(7, 0, s) {
		for _, b := range CountKprimes(3, 0, s-c) {
			a := s - c - b

			if a > 0 && primeFind(a) == 1 {
				result++
			}
		}
	}
	return
}
