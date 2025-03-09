package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var principal, time float64 = 1000, 10
	rate := 5.5

	fmt.Scan(&principal)

	finalValue := principal * math.Pow(1+(rate/100), time)

	finalValueAdjustedInflation := finalValue / math.Pow(1+inflationRate/100, time)

	fmt.Println(finalValue)
	fmt.Println(finalValueAdjustedInflation)

}
