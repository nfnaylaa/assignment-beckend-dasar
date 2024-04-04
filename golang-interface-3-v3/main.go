package main

import (
	"fmt"
	"strconv"
)

type Time struct {
	Hour   int
	Minute int
}

func ChangeToStandartTime(time interface{}) string {

	switch input := time.(type) {
	case string:
		if len(input) != 5 || input[2] != ':' {
			return "Invalid input"
		}
		hour, err := strconv.Atoi(input[:2])
		if err != nil || hour < 0 || hour > 23 {
			return "Invalid input"
		}
		minute, err := strconv.Atoi(input[3:])
		if err != nil || minute < 0 || minute > 59 {
			return "Invalid input"
		}

		return formatTime(hour, minute)
	case []int:
		if len(input) != 2 {
			return "Invalid input"
		}
		if input[0] < 0 || input[0] > 23 || input[1] < 0 || input[1] > 59 {
			return "Invalid Input"
		}
		return formatTime(input[0], input[1])
	case map[string]int:
		hour, hourTrue := input["hour"]
		if !hourTrue || hour < 0 || hour > 23 {
			return "Invalid input"
		}

		minute, minuteTrue := input["minute"]
		if !minuteTrue || minute < 0 || minute > 59 {
			return "Invalid input"
		}
		return formatTime(hour, minute)
	case Time:
		if input.Hour < 0 || input.Hour > 23 || input.Minute < 0 || input.Minute > 59 {
			return "Invalid input"
		}
		return formatTime(input.Hour, input.Minute)
	default:
		return "Invalid input"

	}
}

func formatTime(hour, minute int) string {
	var period string
	if hour >= 12 {
		period = "PM"
		if hour > 12 {
			hour -= 12
		}
	} else {
		period = "AM"
		if hour == 0 {
			hour = 12
		}
	}
	return fmt.Sprintf("%02d:%02d %s", hour, minute, period)
}

func main() {
	fmt.Println(ChangeToStandartTime("16:00"))
	fmt.Println(ChangeToStandartTime([]int{16, 0}))
	fmt.Println(ChangeToStandartTime(map[string]int{"hour": 16, "minute": 0}))
	fmt.Println(ChangeToStandartTime(Time{16, 0}))
}
