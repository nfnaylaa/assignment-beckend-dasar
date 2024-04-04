package main

import (
	"fmt"
	"strings"
)

func CountingLetter(text string) int {

	var unreadableLetters string = "RSTZrstz"
	var count int = 0
	for _, letter := range text {
		if strings.ContainsRune(unreadableLetters, letter) {
			count++
		}
	}
	return count // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Semangat"))
}
