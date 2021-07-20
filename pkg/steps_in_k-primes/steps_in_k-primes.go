package stepsinkprimes

func KPrimesStep(k int, step int, m int, n int) [][]int {
	var NPrimeSlice []int

	for i := m; i <= n; i++ {
		if primeFind(i) == k {
			NPrimeSlice = append(NPrimeSlice, i)
		}
	}

	return stepFind(NPrimeSlice, step)
}

func stepFind(NPrime []int, step int) (result [][]int) {
	for k, current := range NPrime {
		for i := k + 1; i < len(NPrime); i++ {
			if NPrime[i]-current == step {
				result = append(result, []int{current, NPrime[i]})
			}
			if NPrime[i]-current >= step {
				break
			}
		}
	}
	return
}

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
