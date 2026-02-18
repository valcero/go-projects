package main

import "fmt"

func main() {
	process()
}
func process() {
	defer fmt.Println("deferred statement")
	fmt.Println("normal statement")
}
