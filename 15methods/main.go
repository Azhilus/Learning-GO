package main

import "fmt"

func main() {
	fmt.Println("Methods in Go")

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 22}
	fmt.Println(hitesh)
	fmt.Printf("Hitesh details are: %+v\n", hitesh)

	fmt.Printf("Name is %v and email is %v\n", hitesh.Name, hitesh.Email)
	hitesh.GetStatus()
	hitesh.NewMail()

	fmt.Printf("Name is %v and email is %v\n", hitesh.Name, hitesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
