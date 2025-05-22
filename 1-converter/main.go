package main

import (
	"fmt"
)

/*
ТЗ

Финализируем приложение калькулятор. Для этого:
	- Сделать меню с шагами
		- Ввод исходной валюты (подсказываем варианты), если ошибка, заново вводим
		- Ввод числа - если ошибка, заново вводим
		- Ввод целевой валюты (подсказываем варианты) - если ошибка, заново вводим

Выделить функцию ввода/ проверки валюты и числа
После получения всех данных с помощью if / switch вычислений итог и вывести результат.
*/

const USD = 1.0
const EUR = 0.90
const RUB = 81.0

var rates = map[string]float64{
	"USD": USD,
	"EUR": EUR,
	"RUB": RUB,
}

func main() {
	fmt.Println(" __ Калькулятор валют __")
	currenciesFrom, count, currenciesTo := scanUserData()
	result := convert(count, currenciesFrom, currenciesTo)
	fmt.Println("Итого:", result)
}

func convert(amount float64, from string, to string) float64 {
	usdValue := amount / rates[from]
	return usdValue * rates[to]
}

func scanUserData() (string, float64, string) {
	var count float64
	currenciesFrom := scanCurrency("Введите исходную валюту (USD/EUR/RUB): ")

	for {
		fmt.Print("Введите число: ")
		_, err := fmt.Scan(&count)
		if err != nil {
			fmt.Println("Вы ввели некорректное число")
			continue
		}

		break
	}

	currenciesTo := scanCurrency("Введите целевую валюту (USD/EUR/RUB): ")

	for currenciesFrom == currenciesTo {
		println("Валюты не должны совпадать.")
		currenciesTo = scanCurrency("Введите целевую валюту (USD/EUR/RUB): ")
	}

	return currenciesFrom, count, currenciesTo
}

func checkCurrencies(currencies string) bool {
	if currencies == "USD" || currencies == "EUR" || currencies == "RUB" {
		return true
	}

	return false
}

func scanCurrency(prompt string) string {
	var currency string
	for {
		fmt.Print(prompt)
		_, err := fmt.Scan(&currency)

		if err != nil {
			fmt.Println("Вы ввели некорректную валюту, попробуйте еще раз")
			continue
		}

		if checkCurrencies(currency) {
			break
		}

		fmt.Println("Введите корректную валюту")
	}
	return currency
}
