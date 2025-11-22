package main

import "fmt"

func main() {
    const USD_TO_EUR float64 = 0.8686
    const USD_TO_RUB float64 = 79.1
    const EUR_TO_RUB float64 = USD_TO_RUB / USD_TO_EUR

    fmt.Printf("%.2f", EUR_TO_RUB)
}