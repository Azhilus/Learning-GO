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

/*
   Notes for Class 3 - bufio, Comma-Ok Syntax, and Error Handling in GoLang:

   1. The "bufio" package in Go provides buffered I/O operations, enhancing the performance of I/O operations.

   2. In this program, a bufio reader is created using "bufio.NewReader(os.Stdin)" to read input from the console.

   3. The welcome message is displayed to greet the user.

   4. The program prompts the user to enter their name, and the "ReadString('\n')" function is used to read the input
      until the Enter key is pressed.

   5. The comma-ok syntax is utilized to handle multiple return values. "input" receives the user's input as a string,
      and "err" holds any potential errors that may occur during the input reading process.

   6. A greeting message is printed to the console, using the user's input. This showcases the interaction between
      the user and the program.

   7. The "Printf" function is used to print the type of the "input" variable, providing insight into the data type
      of the user's input.

   8. Error handling is performed by checking if the "err" variable is not nil. If "err" is not nil, it indicates
      that an error occurred during the input reading process.

   9. The combination of bufio, comma-ok syntax, and error handling demonstrates robust user input handling,
      preventing unexpected crashes and providing informative messages to users in case of errors.

   10. Understanding how to use bufio for input, comma-ok syntax for handling multiple return values, and implementing
       error handling practices is essential for building reliable and user-friendly Go programs.

   11. Go's simplicity, combined with its effective error handling mechanisms, contributes to the creation of clean
       and robust code in real-world applications.
*/
