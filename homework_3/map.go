package main

import (
	"fmt"
	"time"
)

func main() {
	var weekdays map[string]int
	weekdays = map[string]int{
		"Mon": 1,
		"Tue": 2,
		"Wed": 3,
		"Thu": 4,
		"Fri": 5,
		"Sut": 6,
		"Sun": 7,
	}

	currentWeekday := time.Now().Weekday().String()[:3]

	dayNumber, found := weekdays[currentWeekday]

	if found {
		fmt.Printf("Current day of the week: %d\n", dayNumber)
	} else {
		fmt.Println("Unknown day of the week")
	}

	for weekday := range weekdays {
		fmt.Println(weekday)
	}
}
