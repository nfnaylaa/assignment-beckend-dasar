package main

import (
	"a21hc3NpZ25tZW50/internal"
	"fmt"
	"strconv"
	"strings"
)

func AdvanceCalculator(calculate string) float32 {

	parts := strings.Split(calculate, " ")
	if len(parts) == 0 {
		return 0
	}

	number, err := strconv.ParseFloat(parts[0], 32)
	if err != nil {
		return 0
	}
	newCalc := internal.NewCalculator(float32(number))
	for i := 1; i < len(parts); i += 2 {
		operator := parts[i]
		operand, err := strconv.ParseFloat(parts[i+1], 32)
		if err != nil {
			return 0
		}
		switch operator {
		case "+":
			newCalc.Add(float32(operand))
		case "-":
			newCalc.Subtract(float32(operand))
		case "*":
			newCalc.Multiply(float32(operand))
		case "/":
			newCalc.Divide(float32(operand))
		default:
			return 0
		}
	}

	return newCalc.Result()
}

func main() {
	res := AdvanceCalculator("3 + 4 / 2 ")

	fmt.Println(res)
}
