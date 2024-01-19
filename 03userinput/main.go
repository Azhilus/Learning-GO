package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Welcome message for the user.
	welcome := "Welcome to Go!"
	fmt.Println(welcome)

	// Creating a reader using bufio for reading input from the console.
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")

	// Using comma-ok syntax to handle multiple return values.
	// "input" receives the user's input, and "err" holds any potential errors.
	input, err := reader.ReadString('\n')

	// Printing a greeting message using the user's input.
	fmt.Println("Hello, ", input, "Nice to meet you!")

	// Printing the type of the input variable.
	fmt.Printf("Type of the input is %T\n", input)

	// Checking for errors during input reading.
	// If err is not nil, it indicates an error occurred during input reading.
	fmt.Println("Error: ", err)
}
