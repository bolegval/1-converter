package main

import (
	"errors"
	"fmt"
	"strconv"
)

const USD_TO_EUR float64 = 0.94
const USD_TO_RUB float64 = 92.0
const EUR_TO_RUB float64 = 1 / USD_TO_EUR * USD_TO_RUB

func main() {

	currencyFrom := getUserInputCurrencyFrom()
	currencyTo := getUserInputCurrencyTo(currencyFrom)
	number := getUserInputNumber()
	result := calculate(number, currencyFrom, currencyTo)

	fmt.Print(result)
}

func getUserInputCurrencyFrom() string {
	var currency string
	for {
		fmt.Print("Введите исходную валюту (USD, EUR, RUB): ")
		fmt.Scan(&currency)

		_, err := checkCurrency(currency)

		if err != nil {
			fmt.Println(err)
			continue
		}

		return currency
	}

}

func getUserInputCurrencyTo(currencyFrom string) string {
	avaliableCurrency := []string{"USD", "EUR", "RUB"}
	var inputCurrency string
	var currency string

	for _, value := range avaliableCurrency {
		if value != currencyFrom {
			if len(inputCurrency) > 0 {
				inputCurrency = inputCurrency + ", "
			}
			inputCurrency = inputCurrency + value
		}
	}

	for {
		fmt.Printf("Введите валюту конвертации (%s): ", inputCurrency)
		fmt.Scan(&currency)

		_, err := checkCurrency(currency)

		if err != nil {
			fmt.Println(err)
			continue
		}

		if currency == currencyFrom {
			fmt.Println("Валюта конвертации не должна совпадать с конвертируемой валютой")
			continue
		}

		return currency
	}

	return ""
}

func getUserInputNumber() float64 {
	var value string

	for {
		fmt.Print("Введите число для конвертации в формате х.х: ")
		fmt.Scan(&value)

		number, err := checkNumber(value)

		if err != nil {
			fmt.Println(err)
			continue
		}

		return number
	}

}

func checkCurrency(currency string) (bool, error) {

	if currency == "USD" || currency == "EUR" || currency == "RUB" {
		return true, nil
	}

	return false, errors.New("Введена недоступная валюта. Повторите ввод")
}

func checkNumber(value string) (float64, error) {
	number, err := strconv.ParseFloat(value, 64)

	if err != nil {
		return 0, errors.New("Введено некорректное число. Повторите ввод")
	}

	return number, nil
}

func calculate(value float64, currencyFrom string, currencyTo string) (result float64) {
	var converter float64

	switch {
	case currencyFrom == "USD" && currencyTo == "EUR":
		converter = USD_TO_EUR
	case currencyFrom == "USD" && currencyTo == "RUB":
		converter = USD_TO_RUB

	case currencyFrom == "EUR" && currencyTo == "RUB":
		converter = EUR_TO_RUB
	}

	result = value * converter

	return result
}
