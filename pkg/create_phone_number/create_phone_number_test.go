package createphonenumber_test

import (
	"testing"

	. "github.com/ptsgr/codewars/pkg/create_phone_number"
)

type testCase struct {
	numbers [10]uint
	result  string
}

var testCases = []testCase{
	{[10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, "(123) 456-7890"},
	{[10]uint{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, "(111) 111-1111"},
	{[10]uint{1, 7, 7, 9, 1, 8, 5, 0, 6, 0}, "(177) 918-5060"},
	{[10]uint{4, 1, 2, 9, 8, 4, 1, 5, 7, 6}, "(412) 984-1576"},
	{[10]uint{5, 6, 8, 8, 7, 7, 7, 8, 0, 5}, "(568) 877-7805"},
	{[10]uint{}, "(000) 000-0000"},
}

func TestCreatePhoneNumber(t *testing.T) {
	for _, test := range testCases {
		result := CreatePhoneNumber(test.numbers)
		if result != test.result {
			t.Error(
				"For", test.numbers,
				"expected", test.result,
				"got", result,
			)
		}
	}

}
