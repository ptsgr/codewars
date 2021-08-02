package perimeterofsquaresinarectangle

func Fibonacci(n int) []int {
	var result []int
	for i := 0; i < n; i++ {
		if i < 2 {
			result = append(result, 1)
		} else {
			result = append(result, result[i-1]+result[i-2])
		}

	}
	return result
}

func Sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func Perimeter(n int) int {
	return 4 * Sum(Fibonacci(n+1))
}
