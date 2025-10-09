package main

import "fmt"

func main() {
	const USD_TO_EUR float32 = 0.94
	const USD_TO_RUB float32 = 92.0
	const EUR_TO_RUB float32 = 1 / USD_TO_EUR * USD_TO_RUB
}

func getUserInput() string {
	var currency string
	fmt.Print("Введите значение: ")
	fmt.Scan(&currency)

	return currency
}

func calculate(value float64, currencyFrom string, currencyTo string) (result float64) {
	return result
}
