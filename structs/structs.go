package main

import "fmt"

func main() {
	p := Person{
		fName: "Hamza",
		lName: "Rehman",
		age:   15,
	}
	fmt.Println(p.fName)
	//anonymous structs
	{
		username: "user123",
		email:    "abc@123",
	}
	fmt.Println(user.email)
}

type Person struct {
	fName string
	lName string
	age   int
}
user := struct {
		username string
		email    string
	}