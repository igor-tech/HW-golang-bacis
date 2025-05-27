package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
ТЗ калькулятор

Нужно создать приложение, которое:
	- принимает операцию (AVG - среднее, SUM - сумму, MED - медиану)
	- принимает неограниченное число числе через запятую (2, 10, 9)
	- Разбивает строку чисел по запятым и затем делает расчет в зависимости от операции выводя результат
*/

var funcOperation = map[int]func([]float64) float64{
	1: calcAvg,
	2: calcSum,
	3: calcMedian,
}

func main() {
	fmt.Println("__ Калькулятор __")

	operation := scanOperation()
	numbers := scanNumberChain()
	result := funcOperation[operation](numbers)

	fmt.Printf("Результат: %.2f\n", result)
}

func calcSum(numbers []float64) float64 {
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func calcAvg(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sum := calcSum(numbers)
	return sum / float64(len(numbers))
}

func calcMedian(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	sort.Float64s(numbers)

	if len(numbers)%2 == 0 {
		return (numbers[len(numbers)/2] + numbers[len(numbers)/2-1]) / 2
	}

	return numbers[len(numbers)/2]
}

func scanOperation() int {
	var operation int
	fmt.Println("Выберите операцию производимую с числами из предложенного списка: \n1. AVG - среднее\n2. SUM - сумма\n3. MED - медиана")

	for {
		fmt.Print("нужно выбрать число (1/2/3): ")
		fmt.Scan(&operation)
		if operation < 1 || operation > 3 {
			fmt.Print("Некорректный выбор операции.")
			continue
		}

		break
	}

	return operation
}

func scanNumberChain() []float64 {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Теперь вводите числа через запятую: ")
	userStringInput, err := reader.ReadString('\n')
	var numbersInput = []float64{}

	if err != nil {
		fmt.Println("Ошибка чтения")
		return numbersInput
	}

	userStringInput = strings.TrimSpace(userStringInput)
	userStringInput = strings.ReplaceAll(userStringInput, " ", "")
	parts := strings.Split(userStringInput, ",")

	for _, part := range parts {
		num, err := strconv.ParseFloat(strings.TrimSpace(part), 64)

		if err != nil {
			fmt.Printf("Ошибка в числе '%s': %v\n", part, err)
			continue
		}

		numbersInput = append(numbersInput, float64(num))
	}
	if len(numbersInput) == 0 {
		fmt.Println("Ни одно число не было корректным")
	}

	return numbersInput
}
