package main

import "fmt"

func main() {
	process()
}
func process() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered", r)
		}
	}()
	fmt.Println("Start process")
	panic("something went wrong")
	fmt.Println("End process")
	//use recover with defer
	//use for logging and reporting
}
