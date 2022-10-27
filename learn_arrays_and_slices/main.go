package main

import "fmt"

func main() {

	// Arrays
	var names [3]string
	names[0] = "Keith"
	names[1] = "Kevin"

	// or

	names2 := [3]string{"Keith", "Kevin", "Kayla"}

	fmt.Println("Names array:", names)
	fmt.Println("Names 2 array", names2)

	fmt.Println("Name[2] is empty?:", names[2] == "")

	// Slices a.k.a variable length arrays
	var name_slice []string
	name_slice = append(name_slice, "Keith")
	fmt.Println("Names:", name_slice)
	name_slice = append(name_slice, "Kevin", "Kayla", "Kyle")

	name_slice2 := []string{}
	name_slice2 = append(name_slice2, "Keith")
	fmt.Println("Names2", name_slice2)
	name_slice2 = append(name_slice2, "Kevin", "Kayla", "Kyle")

	fmt.Println("Names:", name_slice)
	fmt.Println("Names2", name_slice2)

	// slice with a minimum number of values
	var name_slice_make = make([]string, 4)
	name_slice_make[0] = "Keith"
	name_slice_make[1] = "Kevin"
	name_slice_make[2] = "Kayla"
	name_slice_make[3] = "Kyle"

	name_slice_make2 := make([]string, 4)
	name_slice_make2[0] = "Keith"
	name_slice_make2[1] = "Kevin"
	name_slice_make2[2] = "Kayla"
	name_slice_make2[3] = "Kyle"

	fmt.Println("Names", name_slice_make)
	fmt.Println("Names2", name_slice_make2)

	name_slice_make = append(name_slice_make, "Jim")

	fmt.Println("Names", name_slice_make)
}
