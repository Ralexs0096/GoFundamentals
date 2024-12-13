package main

import "fmt"

func main() {
	var revenue int
	var expenses int
	var taxRate float64

	outputText("Revenue: ")
	fmt.Scan(&revenue)

	outputText("Expenses: ")
	fmt.Scan(&expenses)

	outputText("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses // earnings before tax
	profit := float64(ebt) * (1 - taxRate/100)
	ratio := ebt / int(profit)

	formattedProfit := fmt.Sprintf("Earnings Before Tax: %.1f", profit)

	fmt.Println(ebt)
	fmt.Println(formattedProfit)
	fmt.Println(ratio)
}

func outputText(text string) {
	fmt.Print(text)
}
