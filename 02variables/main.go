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

/*
   Notes for Class 2 - Data Types and Variable Declaration in GoLang:

   1. Go is a statically-typed language, which means that variable types must be declared explicitly.
      This helps catch errors at compile time.

   2. Constants in Go are declared using the "const" keyword. They must be assigned a value during declaration,
      and once assigned, their values cannot be changed.

   3. Variables are containers for storing data values. They are declared using the "var" keyword, followed
      by the variable name, the type, and an optional initial value.

   4. In Go, there are various built-in primitive data types, including strings, booleans, integers, unsigned integers,
      floating-point numbers, complex numbers, bytes, and runes.

   5. Advanced data types in Go include arrays, slices, maps, structs, and interfaces.

   6. The "fmt" package is used for formatted I/O. The "Println" function is used to print to the console.

   7. Printf is used for formatted printing. The "%T" verb is used to print the type of a variable.

   8. Type inference in Go allows you to declare a variable without explicitly specifying its type. The compiler
      infers the type based on the assigned value.

   9. Short variable declaration (:=) is a concise way of declaring and initializing variables. It can only be
      used inside functions.

   10. Constants, variables, and data types are fundamental concepts in programming, and understanding them is
       crucial for writing effective and bug-free code.

   11. The code provided demonstrates the declaration and initialization of variables and constants with
       various primitive and advanced data types, showcasing Go's syntax and flexibility.
*/
