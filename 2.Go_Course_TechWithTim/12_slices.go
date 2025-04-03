/*
Slices are a reference type
They are more flexible than arrays
They can grow and shrink in size
They can be created from arrays
They can be created using the make function

If you create a slice from an array, the slice will reference the same underlying array.
So, the pointer of the slice will point to the same memory location as the array.
*/

package main

import "fmt"

func main() {

	// Create an array of integers
	// length of the array is 5
	// the capacity of the array is 5
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	fmt.Println(len(arr)) // length of the array is 5
	fmt.Println(cap(arr)) // capacity of the array is 5

	fmt.Println()

	// Create a slice from the array
	var slice = arr[1:4]
	fmt.Println(slice)
	fmt.Println(len(slice)) // length of the slice is 3
	fmt.Println(cap(slice)) // capacity of the slice is 5

	fmt.Println()

	// reslice the slice to a new length and capacity
	slice = slice[1:3] // this will create a new slice with length 2 and capacity 3
	fmt.Println(slice)
	fmt.Println(len(slice)) // length of the slice is 2
	fmt.Println(cap(slice)) // capacity of the slice is 3

	fmt.Println()

	// Create a slice using the make function
	var slice2 = make([]int, 5) // useful in concurrency
	fmt.Println(slice2)

	fmt.Println()

	// Create a slice by not specifying the capacity
	var slice3 []int // this is a nil slice
	fmt.Println(slice3)

	fmt.Println()

	// append to the slice
	slice3Modified := append(slice3, 1)
	fmt.Println(slice3Modified) // this will create a new slice

}
