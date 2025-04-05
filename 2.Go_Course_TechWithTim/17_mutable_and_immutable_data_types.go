/*
	Immutable and Mutable Data Types

Mutable data types: // Data can be changed
	- Slices
	- Arrays
	- Maps
	- Structs
	- Pointers

Immutable data types: // Data cannot be changed
	- Strings
	- Integers
	- Floats
	- Booleans
	- Constants (declared using const) are immutable by definition.
		const pi = 3.14
		pi = 3.15 // This will result in a compile-time error
*/

package main

import "fmt"

func main() {

	// ############################
	// Immutable Data Types
	fmt.Println("Immutable Data Types")

	var x int = 5

	y := x
	y = 10

	fmt.Printf("x: %v, y: %v\n", x, y) // x = 5, y = 10
	// that was an example of immutable data types
	// immutable data types are data types that cannot be changed
	// so why y was different that x?
	// because y is a copy of x
	// so when we change y, x is not affected

	fmt.Println()

	// ############################
	// Mutable Data Types
	fmt.Println("Mutable Data Types")

	var mySlice []int = []int{1, 2, 3, 4, 5}

	mySlice2 := mySlice // means mySlice2 points to the same memory address as mySlice
	mySlice2[0] = 100

	fmt.Printf("mySlice: %v, mySlice2: %v\n", mySlice, mySlice2) // mySlice = [100 2 3 4 5], mySlice2 = [100 2 3 4 5]
}
