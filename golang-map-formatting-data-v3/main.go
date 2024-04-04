package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ChangeOutput(data []string) map[string][]string {
	output := make(map[string][]string)

	for _, datas := range data {
		parts := strings.Split(datas, "-")
		key := parts[0]
		index, _ := strconv.Atoi(parts[1])
		firstOrLast := parts[2]
		values := parts[3]

		if _, found := output[key]; !found {
			output[key] = make([]string, 0)
		}

		if index < len(output[key]) {
			if firstOrLast == "first" {
				output[key][index] = values + " " + output[key][index]
			} else {
				output[key][index] += " " + values
			}
		} else {
			output[key] = append(output[key], values)
		}
	}

	return output
}

func main() {
	data := []string{"account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe", "address-0-first-Jaksel", "address-0-last-Jakarta", "address-1-first-Bandung", "address-1-last-Jabar", "phone-0-first-081234567890", "phone-1-first-081234567891"}
	res := ChangeOutput(data)

	fmt.Println(res)
}
