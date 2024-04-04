package main

import (
	"fmt"
	"strconv"
)

func DateFormat(day, month, year int) string {

	var newDay string
	var newMonth string
	if day < 10 {
		newDay = "0" + strconv.Itoa(day)
	} else {
		newDay = strconv.Itoa(day)
	}

	switch month {
	case 1:
		newMonth = "January"
	case 2:
		newMonth = "February"
	case 3:
		newMonth = "March"
	case 4:
		newMonth = "April"
	case 5:
		newMonth = "May"
	case 6:
		newMonth = "June"
	case 7:
		newMonth = "July"
	case 8:
		newMonth = "August"
	case 9:
		newMonth = "September"
	case 10:
		newMonth = "October"
	case 11:
		newMonth = "November"
	case 12:
		newMonth = "December"
	}

	return newDay + "-" + newMonth + "-" + strconv.Itoa(year) // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(1, 1, 2012))
}
