package reverseorrotate_test

import (
	"testing"

	. "github.com/ptsgr/codewars/pkg/reverse_or_rotate"
)

type testCase struct {
	str    string
	sz     int
	output string
}

var testCases = []testCase{
	{"1234", 0, ""},
	{"", 0, ""},
	{"1234", 5, ""},
	{"733049910872815764", 5, "330479108928157"},
}

func TestRevrot(t *testing.T) {
	for _, test := range testCases {
		result := Revrot(test.str, test.sz)
		if result != test.output {
			t.Error(
				"For", test.str, " sz: ", test.sz,
				"expected", test.output,
				"got", result,
			)
		}
	}
}
