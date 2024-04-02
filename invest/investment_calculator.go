package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	outputText("Investment amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	outputText("Years: ")
	fmt.Scan(&years)

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, float64(years))
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, float64(years))
	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	fmt.Printf("Future value : %v \n", futureValue)
	fmt.Printf("Future value float type : %.1f \n", futureValue)

	// fmt.Sprintf()
	// %.0f , %.1f ... for format float value
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}

// alternative return value
func calculateFutureValues2(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return
}
