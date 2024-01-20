package main

import "fmt"

func main() {
	fmt.Println("Functions in GOLANG")

	greeter()
	greeterTwo()

	result := adder(5, 6)
	fmt.Println("Result is: ", result)

	proResult := proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("Pro Result is: ", proResult)
}

func greeter() {
	fmt.Println("Jai Shree Ram")
}

func greeterTwo() {
	fmt.Println("Radhe Krishna")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}
