package doiphantineequation_test

import (
	"reflect"
	"testing"

	. "github.com/ptsgr/codewars/pkg/doiphantine_equation"
)

type testCase struct {
	n      int
	result [][]int
}

var testCases = []testCase{
	{5, [][]int{{3, 1}}},
	{12, [][]int{{4, 1}}},
	{13, [][]int{{7, 3}}},
	{9005, [][]int{{4503, 2251}, {903, 449}}},
	{9008, [][]int{{1128, 562}}},
	{90002, [][]int{}},
}

func TestSolequa(t *testing.T) {
	for _, test := range testCases {
		result := Solequa(test.n)
		if !reflect.DeepEqual(result, test.result) {
			t.Error(
				"For", test.n,
				"expected", test.result,
				"got", result,
			)
		}
	}

}
