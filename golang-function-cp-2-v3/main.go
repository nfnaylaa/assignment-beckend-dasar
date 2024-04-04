package main

import (
	"fmt"
)

func CountVowelConsonant(str string) (int, int, bool) {
	var vowel = "aiueoAIUEO"
	var countVowel, countConsonant int

	for i := len(str) - 1; i >= 0; i-- {
		if string(str[i]) == " " {
			continue
		}

		isVowel := false
		for j := 0; j < len(vowel); j++ {
			if string(str[i]) == string(vowel[j]) {
				isVowel = true
				break
			}
		}

		if isVowel {
			countVowel++
		} else {
			if (str[i] >= 'a' && str[i] <= 'z') || (str[i] >= 'A' && str[i] <= 'Z') {
				countConsonant++
			}
		}
	}

	if (countVowel == 0 && countConsonant == 0) || countConsonant == 0 || countVowel == 0 {
		return countVowel, countConsonant, true
	}

	return countVowel, countConsonant, false
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountVowelConsonant("Hidup Itu Indah"))
}
