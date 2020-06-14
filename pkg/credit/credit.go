package credit

import (
	"math"
)

func Calculate(amount int64, year int64, percent float64) (monthly, over, total  int64) {
	months := year * 12
	amountCents := amount * 100
	monthlyPercent := percent / 100 / 12

	denominator := math.Pow(1 + monthlyPercent, float64(months)) - 1
	monthly = int64(float64(amountCents) * (monthlyPercent + monthlyPercent / denominator))

	total = monthly * months

	over =  total - amountCents

	return
}







