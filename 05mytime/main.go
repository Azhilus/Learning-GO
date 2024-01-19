package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()

	fmt.Println("Present time is: ", presentTime)

	fmt.Println("Year: ", presentTime.Format("01-02-2006 15:04 Monday"))

	createdDate := time.Date(2020, time.November, 10, 23, 23, 0, 0, time.UTC)

	fmt.Println("Created date is: ", createdDate.Format("01-02-2006 15:04 Monday"))
}
