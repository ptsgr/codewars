package buyingacar

import (
	"math"
)

func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
	priceNew, priceOld := float64(startPriceNew), float64(startPriceOld)
	monthCount, saveMoneyCount := 0, 0
	lossPercent := percentLossByMonth
	if startPriceOld >= startPriceNew {
		return [2]int{0, startPriceOld - startPriceNew}
	}

	for saveMoneyCount+int(priceOld) < int(priceNew) {
		monthCount++
		if monthCount%2 == 0 {
			lossPercent += 0.5
		}
		saveMoneyCount += savingperMonth

		priceOld = priceOld * (100 - lossPercent) / 100
		priceNew = priceNew * (100 - lossPercent) / 100

	}
	return [2]int{monthCount, saveMoneyCount + int(math.Round(priceOld-priceNew))}
}
