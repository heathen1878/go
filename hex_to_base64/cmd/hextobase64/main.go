package main

import (
	"fmt"
	"hextobase64"
)

func main() {
	var hex_value string

	fmt.Printf("Enter hex value to convert to base64: ")

	fmt.Scanln(&hex_value)

	b64Str := hextobase64.HexToBase64(hex_value)
	fmt.Printf("Your base64 value is: %s", b64Str)
}
