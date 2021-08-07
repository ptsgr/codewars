package railfencecipher_test

import (
	"testing"

	. "github.com/ptsgr/codewars/pkg/rail_fence_cipher"
)

type testCase struct {
	decodedStr string
	n          int
	encodedStr string
}

var testCases = []testCase{
	{"WEAREDISCOVEREDFLEEATONCE", 3, "WECRLTEERDSOEEFEAOCAIVDEN"},
	{"Hello, World!", 3, "Hoo!el,Wrdl l"},
	{"Welcom from @ptsgr Codewars repository", 3, "Wor@goarsrecmfo psrCdwr eoioyl mt espt"},
	{"WEAREDISCOVEREDFLEEATONCE", 4, "WIREEEDSEEEACAECVDLTNROFO"},
	{"Hello, World!", 4, "H !e,Wdloollr"},
	{"Welcom from @ptsgr Codewars repository", 4, "W @ apremf prCwreooylormtgoesrstcosd i"},
}

func TestEncode(t *testing.T) {
	for _, test := range testCases {
		result := Encode(test.decodedStr, test.n)
		if test.encodedStr != result {
			t.Error(
				"For", test.decodedStr, "with", test.n, "rails\n",
				"expected:", test.encodedStr,
				"got:", result,
			)
		}
	}

}

func TestDecode(t *testing.T) {
	for _, test := range testCases {
		result := Decode(test.encodedStr, test.n)
		if test.decodedStr != result {
			t.Error(
				"For", test.encodedStr, "with", test.n, "rails\n",
				"expected:", test.decodedStr,
				"got:", result,
			)
		}
	}
}
