package buyingacar_test

import (
	"testing"

	. "github.com/ptsgr/codewars/pkg/buying_a_car"
)

type testCase struct {
	startPriceOld, startPriceNew, savingperMonth int
	percentLossByMonth                           float64
	result                                       [2]int
}

var testCases = []testCase{
	{2000, 8000, 1000, 1.5, [2]int{6, 766}},
	{12000, 8000, 1000, 1.5, [2]int{0, 4000}},
	{8000, 12000, 500, 1.0, [2]int{8, 597}},
	{18000, 32000, 1500, 1.25, [2]int{8, 332}},
	{7500, 32000, 300, 1.55, [2]int{25, 122}},
}

func TestCreatePhoneNumber(t *testing.T) {
	for _, test := range testCases {
		result := NbMonths(test.startPriceOld, test.startPriceNew, test.savingperMonth, test.percentLossByMonth)
		if result != test.result {
			t.Error(
				"For", test,
				"expected", test.result,
				"got", result,
			)
		}
	}

}
