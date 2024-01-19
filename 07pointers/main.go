package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go")

	// Declare a variable 'myNumber' and assign a value.
	myNumber := 23

	// Declare a pointer variable 'ptr' and assign the address of 'myNumber' to it.
	var ptr = &myNumber

	// Print the value stored in 'ptr', which is the memory address of 'myNumber'.
	fmt.Println("Value of myNumber is: ", ptr)

	// Print the value stored at the memory address pointed to by 'ptr'.
	fmt.Println("Value of myNumber is: ", *ptr)

	// Print the memory address of the 'ptr' variable itself.
	fmt.Println("Address of myNumber is: ", &ptr)

	// Multiply the value at the memory address pointed to by 'ptr' by 2.
	*ptr = *ptr * 2

	// Print the updated value of 'myNumber'.
	fmt.Println("New value is: ", myNumber)
}
