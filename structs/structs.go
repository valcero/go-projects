package main

import "fmt"

func main() {
	type Person struct {
		fName string
		lName string
		age   int
	}
	p := Person{
		fName: "Hamza",
		lName: "Rehman",
		age:   15,
	}
	fmt.Println(p.fName)
}
