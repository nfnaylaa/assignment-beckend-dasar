package main

import (
	"fmt"
	"sort"
)

func SchedulableDays(villager [][]int) []int {

	if len(villager) == 0 {
		return []int{}
	}
	for _, villagers := range villager {
		sort.Ints(villagers)
	}

	free := villager[0]

	for _, villagers := range villager[1:] {
		freeDays := []int{}

		i, j := 0, 0
		for i < len(free) && j < len(villagers) {
			if free[i] == villagers[j] {
				freeDays = append(freeDays, free[i])
				i++
				j++
			} else if free[i] < villagers[j] {
				i++
			} else {
				j++
			}

		}

		free = freeDays

	}
	return free // TODO: replace this

}

func main() {
	ruangDesa := [][]int{{7, 12, 19, 22}, {12, 19, 21, 23}, {7, 12, 19}, {12, 19}}
	fmt.Println(SchedulableDays(ruangDesa))
}
