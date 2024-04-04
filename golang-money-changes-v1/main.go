package main

import (
	"fmt"
)

type Product struct {
	Name  string
	Price int
	Tax   int
}

func ExchangeCoin(amount int) []int {

	coins := []int{1000, 500, 200, 100, 50, 20, 10, 5, 1}
	change := []int{}

	for _, coin := range coins {
		for amount >= coin {
			change = append(change, coin)
			amount -= coin
		}
	}

	return change // TODO: replace this
}

func MoneyChanges(amount int, products []Product) []int {

	total := 0
	for _, p := range products {
		total += p.Price + p.Tax
	}

	change := amount - total

	return ExchangeCoin(change) // TODO: replace this
}
func main() {
	amount := 10000
	products := []Product{
		{Name: "Baju", Price: 5000, Tax: 500},
		{Name: "Celana", Price: 3000, Tax: 300},
	}
	result := MoneyChanges(amount, products)
	fmt.Println(result)
}
