package main

import "fmt"

func main() {

	ages := map[string]int{}
	ages["Kevin"] = 18

	if ages["Kevin"] < 67 {
		fmt.Println("Kevin is", ages["Kevin"], "retirement age is 67")
	} else {
		fmt.Println("Kevin is", ages["Kevin"], "and is due to retire")
	}

	switch {
	case ages["Kevin"] < 67:
		fmt.Println("Kevin is", ages["Kevin"], "retirement age is 67")
	default:
		fmt.Println("Kevin is", ages["Kevin"], "and is due to retire")
	}

	switch ages["Kevin"] {
	case 18:
		fmt.Println("Kevin's", ages["Kevin"], "th!")
	}

}
