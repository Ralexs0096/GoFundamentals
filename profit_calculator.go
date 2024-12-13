package main

import "fmt"

func main() {
	var revenue int
	var expenses int
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses // earnings before tax
	profit := float64(ebt) * (1 - taxRate/100)
	ratio := ebt / int(profit)

	formattedProfit := fmt.Sprintf("Earnings Before Tax: %.1f", profit)

	fmt.Println(ebt)
	fmt.Println(formattedProfit)
	fmt.Println(ratio)
}
