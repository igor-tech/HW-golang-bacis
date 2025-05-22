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
	var currenciesFrom string
	var currenciesTo string
	var count float64
	for {
		fmt.Print("Введите исходную валюту (USD/EUR/RUB): ")
		_, err := fmt.Scan(&currenciesFrom)

		if err != nil {
			fmt.Println("Вы ввели некорректную валюту, попробуйте еще раз")
			continue
		}

		if checkCurrencies(currenciesFrom) {
			break
		}

		fmt.Println("Введите корректную валюту")
	}

	for {
		fmt.Print("Введите число: ")
		_, err := fmt.Scan(&count)
		if err != nil {
			fmt.Println("Вы ввели некорректное число")
			continue
		}

		break
	}

	for {
		fmt.Print("Введите целевую валюту (USD/EUR/RUB): ")
		_, err := fmt.Scan(&currenciesTo)

		if err != nil {
			fmt.Println("Вы ввели некорректную валюту, попробуйте еще раз")
			continue
		}
		if currenciesTo == currenciesFrom {
			fmt.Println("Валюты не должны повторятся")
			continue
		}

		if checkCurrencies(currenciesTo) {
			break
		}

		fmt.Println("Введите корректную валюту")
	}
	return currenciesFrom, count, currenciesTo
}

func checkCurrencies(currencies string) bool {
	if currencies == "USD" || currencies == "EUR" || currencies == "RUB" {
		return true
	}

	return false
}
