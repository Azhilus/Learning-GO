// Package main is a special package in Go. It defines an executable program.
package main

// Importing the "fmt" package, which provides functions for formatted I/O.
import "fmt"

// The main function is the entry point of the program.
func main() {
	// Printing "Hello, World!" to the standard output (console).
	fmt.Println("Hello, World!")
}

/*
   Notes for Class 1 - Introduction to GoLang:

   1. Go (or Golang) is a statically-typed, compiled language designed for simplicity, efficiency,
      and readability. It was developed by Google.

   2. "package main" indicates that this file is the start of the executable program.

   3. "import" statement is used to include packages that provide functionality beyond the
      basic language features. Here, "fmt" is imported for formatted I/O.

   4. The "func main()" function is the starting point of the program. Execution begins here.

   5. The "fmt.Println()" function is used to print output to the console. In this case,
      it prints the famous "Hello, World!" message.

   7. Go is a statically-typed language, meaning variable types are declared explicitly.

   8. Go does not support traditional object-oriented features like class-based inheritance,
      but it does support types and methods, making it suitable for both procedural and
      object-oriented programming.

   9. Go emphasizes simplicity and readability. It has a garbage collector for automatic
      memory management.

   10. Go uses goroutines for concurrent programming, making it efficient for handling
       concurrent tasks.

   11. Unlike many other languages, Go does not support method overloading. However,
       it supports function overloading by allowing multiple functions with the same
       name in the same package.

   12. Go has a unique approach to error handling using multiple return values,
       where the last return value is often an error.

   13. Go programs are compiled into machine code, resulting in a single binary executable.
*/
