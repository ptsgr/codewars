package stringsmix_test

import (
	"testing"

	. "github.com/ptsgr/codewars/pkg/StringsMix"
)

type testData struct {
	s1     string
	s2     string
	output string
}

var tests = []testData{
	{"Are they here", "yes, they are here", "2:eeeee/2:yy/=:hh/=:rr"},
	{"uuuuuu", "uuuuuu", "=:uuuuuu"},
	{"looping is fun but dangerous", "less dangerous than coding",
		"1:ooo/1:uuu/2:sss/=:nnn/1:ii/2:aa/2:dd/2:ee/=:gg"},
	{"my&friend&Paul has heavy hats! &", "my friend John has many many friends &",
		"2:nnnnn/1:aaaa/1:hhh/2:mmm/2:yyy/2:dd/2:ff/2:ii/2:rr/=:ee/=:ss"},
	{"mmmmm m nnnnn y&friend&Paul has heavy hats! &", "my frie n d Joh n has ma n y ma n y frie n ds n&",
		"1:mmmmmm/=:nnnnnn/1:aaaa/1:hhh/2:yyy/2:dd/2:ff/2:ii/2:rr/=:ee/=:ss"},
	{"codewars", "codewars", ""},
}

func TestMix(t *testing.T) {
	for _, test := range tests {
		result := Mix(test.s1, test.s2)
		if result != test.output {
			t.Error(
				"For", test.s1, test.s2,
				"expected", test.output,
				"got", result,
			)
		}
	}
}
