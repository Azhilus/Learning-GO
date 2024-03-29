// Package main is a special package in Go. It defines an executable program.
package main

// Importing the "fmt" package, which provides functions for formatted I/O.
import "fmt"

// Name is a constant string with the value "John Doe".
const Name string = "John Doe"

func main() {
	// Variable declaration and initialization with a string value.
	var message string = "Hello, World!"
	fmt.Println(message)
	fmt.Printf("Variable is of type: %T\n", message)

	// Variable declaration and initialization with a boolean value.
	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("Variable is of type: %T\n", isLoggedin)

	// Variable declaration and initialization with an integer value.
	var number int = 10
	fmt.Println(number)
	fmt.Printf("Variable is of type: %T\n", number)

	// Variable declaration and initialization with an unsigned 8-bit integer value.
	var smallNumber uint8 = 255
	fmt.Println(smallNumber)
	fmt.Printf("Variable is of type: %T\n", smallNumber)

	// Variable declaration and initialization with a floating-point value.
	var smallFloat float32 = 255.4555555
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T\n", smallFloat)

	// Complex number data type in Go.
	var complexNumber complex64 = 3 + 4i
	fmt.Println(complexNumber)
	fmt.Printf("Variable is of type: %T\n", complexNumber)

	// Byte and rune data types in Go.
	var char byte = 'A'
	var unicodeChar rune = '☺'
	fmt.Println(char)
	fmt.Printf("Variable is of type: %T\n", char)
	fmt.Println(unicodeChar)
	fmt.Printf("Variable is of type: %T\n", unicodeChar)

	// Array data type in Go.
	var integerArray [3]int = [3]int{1, 2, 3}
	fmt.Println(integerArray)
	fmt.Printf("Variable is of type: %T\n", integerArray)

	// Slice data type in Go.
	var floatSlice []float64 = []float64{1.1, 2.2, 3.3}
	fmt.Println(floatSlice)
	fmt.Printf("Variable is of type: %T\n", floatSlice)

	// Map data type in Go.
	var userMap map[string]int = map[string]int{"Alice": 25, "Bob": 30}
	fmt.Println(userMap)
	fmt.Printf("Variable is of type: %T\n", userMap)

	// Struct data type in Go.
	type Person struct {
		Name string
		Age  int
	}
	var personObject Person = Person{"Charlie", 35}
	fmt.Println(personObject)
	fmt.Printf("Variable is of type: %T\n", personObject)

	// Interface data type in Go.
	var genericInterface interface{} = "This is an interface"
	fmt.Println(genericInterface)
	fmt.Printf("Variable is of type: %T\n", genericInterface)

	// Printing the value of the constant variable Name.
	fmt.Println(Name)
	fmt.Printf("Variable is of type: %T\n", Name)
}
