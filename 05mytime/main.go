package main

import (
	"fmt"
	"time"
)

func main() {
	// Welcome message for the time study in Golang.
	fmt.Println("Welcome to time study of Golang")

	// Getting the current time.
	presentTime := time.Now()

	// Printing the present time.
	fmt.Println("Present time is: ", presentTime)

	// Formatting and printing the present time with specific components (Year, Month, Day, Hour, Minute, Day of the Week).
	fmt.Println("Year: ", presentTime.Format("01-02-2006 15:04 Monday"))

	// Creating a specific date and time (November 10, 2020, 23:23:00 UTC).
	createdDate := time.Date(2020, time.November, 10, 23, 23, 0, 0, time.UTC)

	// Formatting and printing the created date.
	fmt.Println("Created date is: ", createdDate.Format("01-02-2006 15:04 Monday"))
}
