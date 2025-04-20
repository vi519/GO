package main

import "fmt"

func main() {
	fmt.Println("Welcome to struct")

	vineet := User{"Vineet", "vineet@gmail.com", false, 28}
	fmt.Println(vineet)
	fmt.Printf("Name of person is %v and email is %v", vineet.Name, vineet.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
