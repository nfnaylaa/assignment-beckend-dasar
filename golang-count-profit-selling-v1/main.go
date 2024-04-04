package main

import "fmt"

func CountProfit(data [][][2]int) []int {
	if len(data) == 0 {
		return []int{}
	}

	totalProfits := make([]int, len(data[0]))

	for i := 0; i < len(data[0]); i++ {
		totalProfit := 0

		for j := 0; j < len(data); j++ {
			profit := data[j][i][0] - data[j][i][1]
			totalProfit += profit
		}

		totalProfits[i] = totalProfit
	}

	return totalProfits
}

func main() {
	data := [][][2]int{
		{{1000, 800}, {700, 500}},
		{{1000, 800}, {900, 200}},
	}
	fmt.Println(CountProfit(data))
}
