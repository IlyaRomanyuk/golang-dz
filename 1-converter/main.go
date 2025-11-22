package main

import "fmt"

func main() {
    const USD_TO_EUR float64 = 0.8686
    const USD_TO_RUB float64 = 79.1
    const EUR_TO_RUB float64 = USD_TO_RUB / USD_TO_EUR

    value, current, target := getUserInput()

    fmt.Println(value, current, target)
}

func calculateResult(value float64, current string, target string) {

}

func getUserInput() (value float64, current string, target string) {
	fmt.Print("Введи значение: ")
	fmt.Scan(&value)
	fmt.Print("Введи текущую валюту: ")
	fmt.Scan(&current)
    fmt.Print("Введи валюту, в которую хотите конвертировать: ")
	fmt.Scan(&target)

	return
}