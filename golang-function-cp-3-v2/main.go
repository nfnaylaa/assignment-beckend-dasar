package main

import (
	"fmt"
	"strings"
)

func FindShortestName(names string) string {

	nameSplit := strings.FieldsFunc(names, func(r rune) bool {
		return r == ' ' || r == ',' || r == ';'
	})
	shortestName := nameSplit[0]

	for _, name := range nameSplit {
		if len(name) < len(shortestName) || (len(name) == len(shortestName) && name < shortestName) {
			shortestName = name
		}
	}

	return shortestName
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan")) // "Tio"
	fmt.Println(FindShortestName("Budi;Tia;Tio"))                         // "Tia"
}
