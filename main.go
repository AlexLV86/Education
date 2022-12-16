package main

import (
	"fmt"
	"module7/calc"
)

func main() {
	var n1, n2 float64
	var op string
	_, err := fmt.Scanln(&n1)
	if err != nil {
		return
	}
	_, err = fmt.Scanln(&op)
	if err != nil {
		return
	}
	_, err = fmt.Scanln(&n2)
	if err != nil {
		return
	}
	calcul := calc.NewCalculator()
	fmt.Println(calcul.Calculate(n1, n2, op))
}
