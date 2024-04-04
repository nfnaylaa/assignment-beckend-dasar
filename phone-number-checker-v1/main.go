package main

import (
	"fmt"
)

func PhoneNumberChecker(number string, result *string) {

	if len(number) < 10 && number[0:2] != "08" || len(number) < 11 && number[0:3] != "628" {
		*result = "invalid"
	} else {
		if number[0:4] == "0811" || number[0:4] == "0812" || number[0:4] == "0813" || number[0:4] == "0814" || number[0:4] == "0815" || number[0:5] == "62811" || number[0:5] == "62812" || number[0:5] == "62813" || number[0:5] == "62814" || number[0:5] == "62815" {
			*result = "Telkomsel"
		} else if number[0:4] == "0816" || number[0:4] == "0817" || number[0:4] == "0818" || number[0:4] == "0819" || number[0:5] == "62816" || number[0:5] == "62817" || number[0:5] == "62818" || number[0:5] == "62819" {
			*result = "Indosat"
		} else if number[0:4] == "0821" || number[0:4] == "0822" || number[0:4] == "0823" || number[0:5] == "62821" || number[0:5] == "62822" || number[0:5] == "62823" {
			*result = "XL"
		} else if number[0:4] == "0827" || number[0:4] == "0828" || number[0:4] == "0829" || number[0:5] == "62827" || number[0:5] == "62828" || number[0:5] == "62829" {
			*result = "Tri"
		} else if number[0:4] == "0852" || number[0:4] == "0853" || number[0:5] == "62852" || number[0:5] == "62853" {
			*result = "AS"
		} else if number[0:4] == "0881" || number[0:4] == "0882" || number[0:4] == "0883" || number[0:4] == "0884" || number[0:4] == "0885" || number[0:4] == "0886" || number[0:4] == "0887" || number[0:4] == "0888" || number[0:5] == "62881" || number[0:5] == "62882" || number[0:5] == "62883" || number[0:5] == "62884" || number[0:5] == "62885" || number[0:5] == "62886" || number[0:5] == "62887" || number[0:5] == "62888" {
			*result = "Smartfren"
		} else {
			*result = "invalid"
		}
	}
}

func main() {
	// bisa digunakan untuk pengujian test case
	var number = "1234567"
	var result string

	PhoneNumberChecker(number, &result)
	fmt.Println(result)
}
