package main

import (
	"fmt"
	"strings"
)

const USD_TO_EUR float64 = 0.8686
const USD_TO_RUB float64 = 79.1
const EUR_TO_RUB float64 = USD_TO_RUB / USD_TO_EUR

type currenсyMap = map[string]float64

func main() {
    fmt.Println("Добро пожаловать в валютный конвертер!")

    current := inputCurrency("Введите исходную валюту:")
    value := inputAmount("Введите сумму:")
    target := inputCurrency("Введите целевую валюту:")

    var converterMap = currenсyMap{"USD": 1, "EUR": USD_TO_EUR, "RUB": USD_TO_RUB}
    result := calculateResult(current, value, target, &converterMap)

    fmt.Printf("%.2f", result)
}

func inputCurrency(prompt string) string {
    fmt.Println(prompt, "USD", "EUR", "RUB")
    var currenсy string
    
    Loop:
    for {
        fmt.Scan(&currenсy)
        switch strings.ToUpper(currenсy) {
        case "USD", "EUR", "RUB":
            break Loop
        default:
            fmt.Println("Попробуйте еще раз:")
            continue
        }
    }

    return strings.ToUpper(currenсy)
}

func inputAmount(prompt string) float64 {
    fmt.Println(prompt)
    var value float64

    for {
        _, err := fmt.Scan(&value)

        if err != nil {
            fmt.Print("Неверное значение, введите еще раз:")
            var discard string
            fmt.Scanln(&discard)
            continue
        }

        if value <= 0 {
            fmt.Print("Значение должно быть больше 0, введите еще раз:")
            continue
        }
        break
    }

    return value
}

func calculateResult(current string, value float64, target string, converterMap *currenсyMap) float64 {
    inUSD := value / (*converterMap)[current]
    return inUSD * (*converterMap)[target]
}