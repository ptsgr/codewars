package stepsinkprimes_test

import (
	"reflect"
	"testing"

	. "github.com/ptsgr/codewars/pkg/steps_in_k-primes"
)

type testData struct {
	k    int
	step int
	m, n int
	res  [][]int
}

var tests = []testData{
	{10, 8, 2425364, 2425700, [][]int{}},
	{6, 8, 2627371, 2627581, [][]int{{2627408, 2627416}, {2627440, 2627448}, {2627496, 2627504}}},
}

func TestKPrimesStep(t *testing.T) {
	for _, test := range tests {
		result := KPrimesStep(test.k, test.step, test.m, test.n)
		if !reflect.DeepEqual(result, test.res) {
			t.Error(
				"For", test.m, test.n,
				"expected", test.res,
				"got", result,
			)
		}
	}
}
