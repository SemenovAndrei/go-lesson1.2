package main

import (
	"fmt"
	"github.com/i-hit/go-lesson1.2/pkg/credit"
)

func main() {
	amount := 1_000_000
	year := 3
	var percent float64= 20
	monthly, over, total := credit.Calculate(amount, year, percent)

	fmt.Println(monthly, over, total)
}