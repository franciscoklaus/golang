package main

import (
	"fmt"
	"math"
)

func main() {
	investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5 // multiple values with super short notation

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Println(futureValue)
}
