package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// var revenue float64
	// var expense float64
	// var taxRate float64

	// fmt.Print("Revenue: ")
	// fmt.Scan(&revenue)

	// fmt.Print("expense: ")
	// fmt.Scan(&expense)

	// fmt.Print("taxRate: ")
	// fmt.Scan(&taxRate)

	// ebt := revenue - expense
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit

	// fmt.Println(ebt)
	// fmt.Println(profit)
	// fmt.Println(ratio)

	revenue, err1 := getUserInput("Revenue: ")
	expense, err2 := getUserInput("expense: ")
	taxRate, err3 := getUserInput("taxRate: ")
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		// return
		panic(err1)
	}

	ebt, profit, ratio := calculateFinancials(revenue, expense, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT : %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("result.txt", []byte(results), 0644)
}

func calculateFinancials(revenue, expense, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expense
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("User input must be positive number.")
	}
	return userInput, nil
}
