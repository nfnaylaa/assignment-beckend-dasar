package main

import "fmt"

func GetPredicate(math, science, english, indonesia int) string {

	// - Jika nilai yang didapat adalah 100, murid mendapat predikat 'Sempurna'
	// - Jika nilai yang didapat 90 keatas, murid mendapat predikat 'Sangat Baik'
	// - Jika nilai yang didapat 80 keatas, murid mendapat predikat 'Baik'
	// - Jika nilai yang didapat 70 keatas, murid mendapat predikat 'Cukup'
	// - Jika nilai yang didapat 60 keatas, murid mendapat predikat 'Kurang'
	// - Jika nilai yang didapat kurang dari 60, murid mendapat predikat 'Sangat kurang'
	var predicate string

	if (math + science + english + indonesia)/4 == 100{
		predicate = "Sempurna"
	}else if (math + science + english + indonesia)/4 >= 90{
		predicate = "Sangat Baik"
	}else if (math + science + english + indonesia)/4 >=80{
		predicate = "Baik"
	}else if (math + science + english + indonesia)/4 >=70{
		predicate = "Cukup"
	}else if(math + science + english + indonesia)/4 >= 60{
		predicate = "Kurang"
	}else{
		predicate = "Sangat kurang"
	}
		return predicate // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetPredicate(100, 100, 100, 100))
}
