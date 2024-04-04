package main

import "fmt"

func SchedulableDays(date1 []int, date2 []int) []int {

	free := []int{}
	i, j := 0, 0
	for i < len(date1) && j < len(date2) {
		if date1[i] == date2[j] {
			free = append(free, date1[i])
			i++
			j++
		} else if date1[i] < date2[j] {
			i++
		} else {
			j++
		}
	}

	return free // TODO: replace this
}
func main() {
	imam := []int{11, 12, 13, 14, 15}
	permana := []int{5, 10, 12, 13, 20, 21}
	fmt.Println(SchedulableDays(imam, permana))
}
