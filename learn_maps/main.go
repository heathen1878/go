package main

import "fmt"

func main() {

	var birthdays = map[string]string{} //[key_type]value_type

	birthdays["Keith"] = "06/02/1990"

	fmt.Println(birthdays)

	birthdays2 := map[string]string{
		"Keith": "06/02/1990",
	}

	fmt.Println(birthdays2)

	// removing items
	delete(birthdays, "Keith")

}
