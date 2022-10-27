package main

import (
	"fmt"
)

func main() {
	ages := map[string]int{}
	ages["Kevin"] = 11
	ages["Keith"] = 28
	ages["James"] = 67
	ages["Michael"] = 18
	ages["Leigha"] = 16

	// sequencing loop
	for name, age := range ages {
		switch age {
		case 1, 2, 3, 5, 7, 9, 11, 13, 17, 19:
			fmt.Println(fmt.Sprintf("%s's age is a prime number", name))
		case 16:
			fmt.Println(fmt.Sprintf("%s can now drive!", name))
		case 18:
			fmt.Println(fmt.Sprintf("%s can now vote!", name))
		case 67:
			fmt.Println(fmt.Sprintf("%s can now retire", name))
		default:
			fmt.Println(fmt.Sprintf("There is nothing special about %s's current age.", name))
		}
	}

	// for loop
	for i := 1; i <= 10; i++ {
		fmt.Println("We're counting", i)
	}

	a := 0
	for a <= 10 {
		fmt.Println("We're counting", a)
		a++ // Increment the a variable on each loop
	}

	b := 0
	for b <= 10 {
		if b%2 == 0 { //get the remainder of 'b' to work out whether its even or not.
			fmt.Println(b, "is even")
			b++
			continue
		} else if b == 5 {
			fmt.Println(b, "is 5")
			fmt.Println("exiting")
			break
		}

		fmt.Println(b, "is odd")
		b++
	}

}
