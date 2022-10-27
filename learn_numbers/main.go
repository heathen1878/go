package main

import (
	"fmt"
	"math"
)

func main() {
	// Basic calculations
	fmt.Println("Addition", 1+3)
	fmt.Println("Subtraction", 1-3)
	fmt.Println("Multiplication", 1*3)
	fmt.Println("Division", 4/3)    // Integer - always returns an integer
	fmt.Println("Division", 10/2)   // Integer - always returns an integer
	fmt.Println("Division", 4.0/3)  // float - always returns a float
	fmt.Println("Division", 10.0/2) // float - always returns a float

	fmt.Println("Exponents", math.Pow(20.1, 3.2))
}
