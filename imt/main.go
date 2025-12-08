package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("-- Калькулятор имт --")

	for {
		weight, height := getUserInput()
		imt := calculateIMT(weight, height)
		outputResult(imt)

		isRepeate := checkRepeatCalculation()

		if !isRepeate {
			break
		}
	}
}

func calculateIMT(weight, height float64) float64 {
	const IMTPower = 2
	IMT := weight / math.Pow(height/100, IMTPower)
	return IMT
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.1f", imt)
	fmt.Println(result)
	switch {
	case imt < 16:
		fmt.Println("У вас сильный дефицит массы тела")
	case imt < 18.5:
		fmt.Println("У вас дефицит массы тела")
	case imt < 25:
		fmt.Println("У вас нормальный вес")
	case imt < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас степень ожирения")
	}
}

func getUserInput() (weight float64, height float64) {
	fmt.Print("Введи рост в сантиметрах:")
	fmt.Scan(&height)
	fmt.Print("Введи вес:")
	fmt.Scan(&weight)

	return
}

func checkRepeatCalculation() bool {
	var repeat string
	fmt.Print("Расчитать имт еще раз (y/n): ")
	fmt.Scan(&repeat)

	if repeat == "y" || repeat == "Y" {
		return true
	}

	return false
}
