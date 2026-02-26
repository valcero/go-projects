package main

func main() {
	type geometry interface {
		area() float64
		perim() float64
	}
}
