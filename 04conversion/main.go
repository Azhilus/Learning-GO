package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Welcome message for the user.
	fmt.Println("Welcome to our pizza shop")
	fmt.Println("Please rate our pizza between 1 and 5: ")

	// Creating a reader using bufio for reading input from the console.
	reader := bufio.NewReader(os.Stdin)

	// Reading user input until the Enter key is pressed.
	input, _ := reader.ReadString('\n')

	// Printing a thank-you message with the user's input.
	fmt.Println("Thanks for rating, ", input)

	// Removing leading and trailing whitespaces, then converting the input to a float64.
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// Checking for errors during the conversion process.
	if err != nil {
		// If an error occurs, printing the error message.
		fmt.Println("Error:", err)
	} else {
		// If no error, adding 1 to the user's rating and printing the result.
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}
