package main

import (
	"dom/message"
	"fmt"
)

func main() {
	m := message.Greeting("Keith", "Hello")
	fmt.Println(m)
}
