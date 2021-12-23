package main

import "fmt"

func weekDay(day int) string {
	switch day {
	case 1:
		return "sunday"
	case 2:
		return "monday"
	case 3:
		return "tuesday"
	case 4:
		return "wednesday"
	case 5:
		return "thursday"
	case 6:
		return "friday"
	case 7:
		return "saturday"
	default:
		return "invalid"
	}
}

func weekDay2(day int) string {
	var weekday string
	switch {
	case day == 1:
		weekday = "sunday"

		// this will force the next condition to be executed, without the check condition
		fallthrough
	case day == 2:
		weekday = "monday"
	case day == 3:
		weekday = "tuesday"
	case day == 4:
		weekday = "wednesday"
	case day == 5:
		weekday = "thursday"
	case day == 6:
		weekday = "friday"
	case day == 7:
		weekday = "saturday"
	default:
		weekday = "invalid"
	}

	return weekday
}

func main() {
	fmt.Println(weekDay(5)) // thursday
	fmt.Println(weekDay(7)) // saturday

	fmt.Println(weekDay2(1)) // monday
}
