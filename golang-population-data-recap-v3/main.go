package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]interface{} {

	var population []map[string]interface{}

	if len(data) == 0 {
		return []map[string]interface{}{}
	}

	for _, datas := range data {
		parts := strings.Split(datas, ";")
		person := make(map[string]interface{})
		person["name"] = parts[0]
		age, _ := strconv.Atoi(parts[1])
		person["age"] = age
		person["address"] = parts[2]
		if len(parts[3]) > 0 {
			height, _ := strconv.ParseFloat(parts[3], 64)
			person["height"] = height
		}

		if len(parts[4]) > 0 {
			isMarried, _ := strconv.ParseBool(parts[4])
			person["isMarried"] = isMarried
		}

		population = append(population, person)

	}

	return population // TODO: replace this
}

func main() {
	data := []string{"Budi;23;Jakarta;;", "Joko;30;Bandung;;true", "Susi;25;Bogor;165.42;"}
	population := PopulationData(data)

	for _, person := range population {
		fmt.Println(person)
	}
}
