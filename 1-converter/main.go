package main

import "fmt"

func main() {
	const USDToEUR = 0.90
	const USDToRUB = 81
	const EURToRUB = USDToRUB / USDToEUR
}

func calculateCurrencies(count int, c1, c2 float64) float64 {
	//
	return c1 + c2
}

func getUserInput() {
	fmt.Print("Enter your currencies: ")
	//fmt.Scan()
}
