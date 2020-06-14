package credit

import (
	"fmt"
	"math"
)

func Calculate(amount int, year int, percent float64) (int64, int64, int64) {
	months := float64(year * 12)
	amountCents := float64(amount * 100)
	monthlyPercent := toFixed(percent / 100 / 12, 4)

	denominator := math.Pow(1 + monthlyPercent, months) - 1
	monthly := amountCents * (monthlyPercent + monthlyPercent / denominator)

	total := monthly * months

	over :=  total - amountCents

	fmt.Println(monthlyPercent, denominator)

	return int64(monthly), int64(over), int64(total)
}

func toFixed(number, point float64) float64 {
	multy := math.Pow(10,(point + 1))
	numberInt := int64(number * multy)
	if numberInt % 10 >= 5 {
		numberInt = numberInt / 10 + 1
	} else {
		numberInt =  numberInt / 10
	}
	number = float64(numberInt) / math.Pow(10, point)
	return number
}






