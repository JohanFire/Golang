/*

Pointers and Derefrence operator (* and &)

& stands for get the pointer = address of the variable
* stands for dereference the pointer

*/

package main

import "fmt"

// change_value_with_pointer is a function that takes a pointer of a string to modify its value
func change_value_with_pointer(str *string) {
	*str = "Modified !"
}

func main() {
	var num int = 10

	fmt.Println("num 	= ", num)
	fmt.Println("&num 	= ", &num) // address of the variable num

	fmt.Println()

	numPointer := &num // numPointer is a pointer to the variable num
	fmt.Println("numPointer 	= &num")
	fmt.Println("numPointer 	= ", numPointer)  // address of the variable num
	fmt.Println("*numPointer = ", *numPointer) // value of the variable num

	fmt.Println()

	// Example using functions
	var stringToChange string = "Tristán"

	fmt.Println(stringToChange)
	change_value_with_pointer(&stringToChange) // pass the address of the variable stringToChange
	fmt.Println(stringToChange)

	fmt.Println()

	// ##########################################
	// example of something weird, a pointer to a pointer
	var myVar string = "This is my var"
	var myPointer *string = &myVar

	fmt.Println("myVar = ", myVar)
	fmt.Println("myPointer = ", myPointer)
	fmt.Println("&myPointer = ", &myPointer) // address of the pointer
	fmt.Println("*myPointer = ", *myPointer) // dereference the pointer to get the value of myVar

}
