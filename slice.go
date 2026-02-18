// slice.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sliceMain() {
	slice := make([]int, 0, 3)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter an integer or 'X' to exit: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if strings.EqualFold(input, "X") {
			fmt.Println("Exiting program.")
			break
		}
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer or 'X' to exit.")
			continue
		}
		slice = append(slice, num)

		sort.Ints(slice)

		fmt.Println("Sorted Slice:", slice)
	}
}
