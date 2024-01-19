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

/*
   Class 4: Type Conversions and Error Handling in GoLang

   1. Welcome Message: The program begins with a welcoming message,
      inviting the user to rate the pizza.

   2. Creating a Reader: A bufio reader is created to facilitate reading input from the console.
      This is a common practice for interactive command-line applications.

   3. User Input: The user is prompted to rate the pizza,
      and their input is read until the Enter key is pressed.

   4. Thank-You Message: A message is printed to thank the user for their rating,
      displaying the entered input.

   5. Type Conversion: The user's input, initially stored as a string,
      needs to be converted to a numerical format for further processing.
      strconv.ParseFloat is used for this purpose.
      The input is trimmed of leading and trailing whitespaces using strings.TrimSpace before the conversion.

   6. Error Handling in Conversion: The program checks for errors during the conversion process.
      If an error occurs, it is captured in the err variable.
      In this case, the error message is printed to the console.

   7. Handling Errors: Effective error handling is crucial for robust programs.
      In case of an error, the program gracefully handles the situation by providing feedback to the user about the encountered issue.
      This improves the user experience and helps developers identify and resolve problems.

   8. Successful Conversion: If the conversion is successful (no error),
      the program proceeds to add 1 to the user's rating and prints the result.
      This demonstrates how to handle both successful and unsuccessful conversions,
      ensuring the program can gracefully recover from unexpected input.

   9. Conclusion: Understanding type conversions and implementing error handling practices is fundamental
      for building reliable and user-friendly applications.
      In this example, the code showcases the importance of these concepts in the context of user input for rating a pizza.
      The ability to handle various scenarios, including unexpected user input, contributes to the overall resilience of the program.

   10. Next Steps: As developers advance, they can explore more sophisticated error handling strategies,
       such as creating custom error types, logging errors, and incorporating user prompts for correction in case of input errors.
       These practices enhance the robustness and user-friendliness of the application.

   11. Overall Significance: The principles demonstrated in this code snippet are applicable across a wide range of scenarios,
       emphasizing the foundational concepts of type conversions and error handling in GoLang.
       These concepts are crucial for developing applications that are both reliable and user-centric.
*/
