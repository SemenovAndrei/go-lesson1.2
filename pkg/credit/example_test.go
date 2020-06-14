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
	//3716358 33788888 133788888
	//371635 3378860 13378860
	//36506 114216 1314216
}