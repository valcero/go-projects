package main

import "fmt"

func main() {
	defer fmt.Println("Deferred 1")
	defer fmt.Println("deferred 2")
	panicProcess(10)
	panicProcess(-2)
}

func panicProcess(i int) {
	if i < 0 {
		panic("int cant be negative")
	}
	fmt.Println("processing number", i)
}
