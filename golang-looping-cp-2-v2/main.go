package main

import "fmt"

// hello World => d_l_r_o_W o_l_l_e_H
func ReverseString(str string) string {

	var reversed string
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}

	var result string
	for i := 0; i < len(reversed); i++ {
		if (i+1 < len(reversed) && reversed[i+1] == ' ') || reversed[i] == ' ' || i == len(reversed)-1 {
			result += string(reversed[i])
		} else {
			result += string(reversed[i]) + "_"
		}

	}

	return result
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseString("Hello World"))
}
