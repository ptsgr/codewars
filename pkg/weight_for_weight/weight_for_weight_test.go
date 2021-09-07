package weightforweight_test

import (
	"testing"

	. "github.com/ptsgr/codewars/pkg/weight_for_weight"
)

type testCase struct {
	input  string
	output string
}

var testCases = []testCase{
	{"", ""},
	{"103 123 4444 99 2000", "2000 103 123 4444 99"},
	{"2000 10003 1234000 44444444 9999 11 11 22 123", "11 11 2000 10003 22 123 1234000 44444444 9999"},
}

func TestMix(t *testing.T) {
	for _, test := range testCases {
		result := OrderWeight(test.input)
		if result != test.output {
			t.Error(
				"For", test.input,
				"expected", test.output,
				"got", result,
			)
		}
	}
}
