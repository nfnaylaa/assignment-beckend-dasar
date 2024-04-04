package main

import (
	"fmt"
	"strconv"
)

func BiggestPairNumber(numbers int) int {
	numStr := strconv.Itoa(numbers)
	maxSum := 0
	maxPair := ""

	for i := 0; i < len(numStr)-1; i++ {
		pairSum := int(numStr[i]-'0') + int(numStr[i+1]-'0')
		if pairSum > maxSum {
			maxSum = pairSum
			maxPair = string(numStr[i]) + string(numStr[i+1])
		}
	}

	maxPairInt, _ := strconv.Atoi(maxPair)

	return maxPairInt
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BiggestPairNumber(89083278))
}
