package main

import (
	"fmt"
)

func FindMin(nums ...int) int {
	minNumber := nums[0]
	for _, num := range nums {
		if num < minNumber {
			minNumber = num
		}
	}
	return minNumber // TODO: replace this
}

func FindMax(nums ...int) int {

	maxNumber := nums[0]
	for _, num := range nums {
		if num > maxNumber {
			maxNumber = num
		}
	}
	return maxNumber // TODO: replace this
}

func SumMinMax(nums ...int) int {

	max := FindMax(nums...)
	min := FindMin(nums...)

	return max + min // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(SumMinMax(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
