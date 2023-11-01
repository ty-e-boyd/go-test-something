package main

import (
	"fmt"
	calculator "testing-go/bin"
)

func main() {
	sum := calculator.Add([]int{1, 2, 3})

	fmt.Printf("Sum of numbers is %v\n", sum)
}
