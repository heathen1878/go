package main

import "fmt"

func main() {

	// Intepteted string literal
	fmt.Println("Simple string")

	// raw string literal
	fmt.Println(`
This is a multiline \n
string.
	`)

	// Intepreted string literals
	fmt.Println("😀")
	fmt.Println("\u2272")

	// Gets the numeric value of the character - rune
	fmt.Println('L')

}
