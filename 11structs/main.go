package main

import "fmt"

func main() {
	fmt.Println("Structs in Go")
	//no inheritance in golang; No super or parent

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 22}
	fmt.Println(hitesh)
	fmt.Printf("Hitesh details are: %+v\n", hitesh)

	fmt.Printf("Name is %v and email is %v\n", hitesh.Name, hitesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
