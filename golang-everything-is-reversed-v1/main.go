package main

import "fmt"

func ReverseData(arr [5]int) [5]int {
	reversed := [5]int{}

	for i := 0; i < len(arr); i++ {
		num := arr[len(arr)-1-i]

		reversed[i] = 0
		for num > 0 {
			digit := num % 10
			reversed[i] = reversed[i]*10 + digit
			num /= 10
		}
	}

	return reversed
}

func main() {
	arr := [5]int{123, 456, 11, 1, 2}
	fmt.Println(ReverseData(arr))
}
