package main

import "fmt"

func BMICalculator(gender string, height int) float64 {

	var count float64

	if gender == "laki-laki" {
		count = (float64(height)-100) - ((float64(height)-100)*0.1)
	} else if gender == "perempuan" {
		count = (float64(height)-100) - ((float64(height)-100)*0.15)
	}
	return count


}

func main() {
	fmt.Println(BMICalculator("laki-laki", 170))
	fmt.Println(BMICalculator("perempuan", 165))
}
