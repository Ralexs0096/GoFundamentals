package main

import (
	"fmt"
	"math"
)

func oldMain() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Introduce the Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount *
		math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+expectedReturnRate/100, inflationRate)

	fmt.Print(futureValue)
	fmt.Print(futureRealValue)
}
