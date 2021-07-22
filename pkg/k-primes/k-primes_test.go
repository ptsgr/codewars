package kprimes_test

import (
	"reflect"
	"testing"

	. "github.com/ptsgr/codewars/pkg/k-primes"
)

type puzzleTest struct {
	s, result int
}

var puzzleTests = []puzzleTest{
	{138, 1},
	{143, 2},
}

func TestPuzzle(t *testing.T) {
	for _, test := range puzzleTests {
		result := Puzzle(test.s)
		if test.result != result {
			t.Error(
				"For", test.s,
				"expected", test.result,
				"got", result,
			)
		}
	}

}

type countKprimesTest struct {
	k, start, nd int
	result       []int
}

var countKprimesTests = []countKprimesTest{
	{5, 1000, 1100, []int{1020, 1026, 1032, 1044, 1050, 1053, 1064, 1072, 1092, 1100}},
	{5, 500, 600, []int{500, 520, 552, 567, 588, 592, 594}},
	{12, 100000, 100100, nil},
}

func TestCountKprimes(t *testing.T) {
	for _, test := range countKprimesTests {
		result := CountKprimes(test.k, test.start, test.nd)
		if !reflect.DeepEqual(result, test.result) {
			t.Error(
				"For", test.k,
				"in interval:", test.start, "to", test.nd,
				"expected", test.result,
				"got", result,
			)
		}
	}

}
