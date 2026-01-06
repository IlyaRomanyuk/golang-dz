package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type calculationFunc = func(numbers []int64) float64
type mapCalculationFunc = map[string]calculationFunc

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Необходимо ввести 2 аргумента в виде строк!")
		os.Exit(1)
	}
	var operation = os.Args[1]
	var inputNumbers = os.Args[2]

	parts := strings.Split(inputNumbers, ",")

	var numbers []int64

	for _, part := range parts {
		num, err := strconv.ParseInt(strings.TrimSpace(part), 10, 64)

		if err != nil {
			fmt.Println("Ошибка преобразования:", err)
			continue
		}

		numbers = append(numbers, num)
	}

	calcFuncs := mapCalculationFunc{
		"sum": calculateSum,
		"avg": calculateAvg,
		"med": calculateMedian,
	}

	var cleanOperation string = strings.TrimSpace(strings.ToLower(operation))

	actionFunc := calcFuncs[cleanOperation]

	if actionFunc == nil {
		fmt.Println("не введена оперция вычисления")
	}

	res := actionFunc(numbers)
	fmt.Println(res)
}

func calculateSum(numbers []int64) (sum float64) {
	for _, value := range numbers {
		sum += float64(value)
	}
	return
}

func calculateAvg(numbers []int64) (avg float64) {
	var sum float64
	var len float64 = float64(len(numbers))

	for _, value := range numbers {
		sum += float64(value)

	}

	avg = sum / len
	return
}

func calculateMedian(numbers []int64) (median float64) {
	slice := make([]int64, len(numbers))

	copy(slice, numbers)

	len := len(slice)

	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})

	if len == 0 {
		return 0
	}

	if len%2 != 0 {
		return float64(slice[len/2])
	} else {
		mid1 := float64(slice[len/2-1])
		mid2 := float64(slice[len/2])
		return (mid1 + mid2) / 2
	}
}
