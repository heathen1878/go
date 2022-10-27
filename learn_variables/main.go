package main

import "fmt"

func main() {
	var i int = 0
	var s string = ""
	var ok = true // variable type is automatically inferred by the assigned value.

	s = "Hello"
	i = 16

	if ok {
		fmt.Println(s, "there!, today I am", i)
	}

}
