package perimeterofsquaresinarectangle_test

import (
	"testing"

	. "github.com/ptsgr/codewars/pkg/perimeter_of_squares_in_a_rectangle"
)

type testCase struct {
	n      int
	result int
}

var testCases = []testCase{
	{5, 80},
	{7, 216},
	{20, 114624},
	{30, 14098308},
}

func TestPerimeter(t *testing.T) {
	for _, test := range testCases {
		result := Perimeter(test.n)
		if test.result != result {
			t.Error(
				"For", test.n,
				"expected", test.result,
				"got", result,
			)
		}
	}

}
