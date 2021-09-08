package ismyfriendcheating_test

import (
	"reflect"
	"testing"

	. "github.com/ptsgr/codewars/pkg/is_my_friend_cheating"
)

type testCase struct {
	m   uint64
	res [][2]uint64
}

var nullResult [][]int

var testCases = []testCase{
	{26, [][2]uint64{{15, 21}, {21, 15}}},
	{100, nil},
	{101, [][2]uint64{{55, 91}, {91, 55}}},
	{102, [][2]uint64{{70, 73}, {73, 70}}},
	{110, [][2]uint64{{70, 85}, {85, 70}}},
}

func TestKPrimesStep(t *testing.T) {
	for _, test := range testCases {
		result := RemovNb(test.m)
		if !reflect.DeepEqual(result, test.res) {
			t.Error(
				"For", test.m,
				"expected", test.res,
				"got", result,
			)
		}
	}
}
