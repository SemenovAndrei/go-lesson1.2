package credit_test

import (
	"fmt"
	"github.com/i-hit/go-lesson1.2/pkg/credit"
)

func ExampleCalculate() {
	fmt.Println(credit.Calculate(1_000_000, 3, 20))
	fmt.Println(credit.Calculate(1_000_00, 3, 20))
	fmt.Println(credit.Calculate(12_000, 3, 6))
	//Output:
	//3718396 33862283 133862283
	//371839 3386228 13386228
	//36506 114227 1314227
}