package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to Go!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")

	// comma ok || err err

	input, err := reader.ReadString('\n')
	fmt.Println("Hello, ", input, "Nice to meet you!")
	fmt.Printf("Type of the input is %T\n", input)
	fmt.Println("Error: ", err)
}
