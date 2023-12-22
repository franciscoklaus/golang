package main

import (
	"fmt"
	"math"
)

func main() {
	//investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5 // multiple values with super short notation

	const inflationRate = 6.5
	var investmentAmount float64 = 1000
	years := 10.0
	expectedReturnRate := 5.5

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
