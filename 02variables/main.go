package main

import "fmt"

const Name string = "John Doe"

func main() {
	var message string = "Hello, World!"
	fmt.Println(message)
	fmt.Printf("Variable is of type: %T\n", message)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("Variable is of type: %T\n", isLoggedin)

	var number int = 10
	fmt.Println(number)
	fmt.Printf("Variable is of type: %T\n", number)

	var smallNumber uint8 = 255
	fmt.Println(smallNumber)
	fmt.Printf("Variable is of type: %T\n", smallNumber)

	var smallFloat float32 = 255.4555555
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T\n", smallFloat)

	var wesite = "https://www.google.com"
	fmt.Println(wesite)
	fmt.Printf("Variable is of type: %T\n", wesite)

	numberofUsers := 300000.00
	fmt.Println(numberofUsers)
	fmt.Printf("Variable is of type: %T\n", numberofUsers)

	fmt.Println(Name)
	fmt.Printf("Variable is of type: %T\n", Name)
}
