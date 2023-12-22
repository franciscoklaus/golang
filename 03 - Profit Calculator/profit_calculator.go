package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100.0)
	ratio := earningsBeforeTax / profit

	fmt.Println("Earnings Before Tax: ", earningsBeforeTax)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)
}
