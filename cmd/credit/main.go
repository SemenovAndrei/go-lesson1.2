package main

import (
	"fmt"
	"github.com/i-hit/go-lesson1.2/pkg/credit"
)

func main() {
	var amount int64 = 1_000_000
	var year int64 = 3
	var percent float64 = 20
	monthly, over, total := credit.Calculate(amount, year, percent)

	fmt.Println(monthly, over, total)
}