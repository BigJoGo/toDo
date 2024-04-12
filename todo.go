package main

import (
	"fmt"
	"time"
)

func getCurrentDate() string {
	currentDate := time.Now()
	formatteDate :=
		currentDate.Format("2006-01-02")
	return formatteDate
}

func main() {

	var Mission string
	var Day string
	var Times float64

	for {
		fmt.Println("Add to plan:")
		fmt.Scan(&Mission)

		fmt.Println("Day of the week:")
		fmt.Scan(&Day)

		fmt.Println("Time of creation:")
		fmt.Scan(&Times)

		var toDo = make(map[string]interface{})
		toDo["Add to plan"] = Mission
		toDo["Day of the week:"] = Day
		toDo["Allarm on time:"] = Times
		toDo["Time:"] = getCurrentDate()

		fmt.Printf("You plan: %v \n", toDo)
	}
}
