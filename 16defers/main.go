package main

import "fmt"

func main() {
	defer fmt.Println("This is printed at the end")
	defer fmt.Println("Is this printed at the middle?")
	defer fmt.Println("Is this printed at the first?")
	fmt.Println("Defer in Go")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
