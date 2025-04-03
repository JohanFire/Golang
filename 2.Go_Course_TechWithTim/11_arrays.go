/*
An array is a collection of items stored at contiguous memory locations.

In Go, arrays are fixed in size.
This means that once you create an array, you cannot change its size.
*/

package main

import "fmt"

func main() {
	var arr []string = []string{"hello", "world", "!"}
	fmt.Println(arr)                              // [hello world !]
	fmt.Println(arr[1])                           // world
	fmt.Printf("arr max size is %d \n", cap(arr)) // The array arr has a max size of 3.
	// Automatically created by Go depending on the number of elements in the array

	fmt.Println()

	// Create an array of integers with 5 elements
	var arr2 [8]int
	fmt.Println(arr2)                                // [1 2 3 4 5]
	fmt.Printf("arr2 has %d elements \n", len(arr2)) // The array arr2 has 4 elements
	fmt.Printf("arr2 max size is %d \n", cap(arr2))  // The array arr2 has a max size of 5

	fmt.Println()

	// Create an array of integers with 5 elements and assign values to it
	var arr3 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)                                // [1 2 3 4 5]
	fmt.Printf("arr3 has %d elements \n", len(arr3)) // The array arr3 has 5 elements
	fmt.Printf("arr3 max size is %d \n", cap(arr3))  // The array arr3 has a max size of 5

	fmt.Println()

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i]) // 1 2 3 4 5
	}

	fmt.Println()

	// Create an array of integers
	var arr4 []int
	fmt.Println(arr4)                                // []
	fmt.Printf("arr4 has %d elements \n", len(arr4)) // The array arr4 has 0 elements
	fmt.Printf("arr4 max size is %d \n", cap(arr4))  // The array arr4 has a max size of 0

	fmt.Println()

	// Create an array of integers with 2 dimensions
	// var array2D [mainArr][subArr]int = [mainArr][subArr]int{{1, 2, 3}, {4, 5, 6}}
	var array2D [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(array2D)                                              // [[1 2 3] [4 5 6]]
	fmt.Printf("The value of array2D[0][2] is: %d \n", array2D[0][2]) // 3
	fmt.Printf("array2D has %d elements \n", len(array2D))            // The array array2D has 2 elements
	fmt.Printf("array2D max size is %d \n", cap(array2D))             // The array array2D has a max size of 2

	fmt.Println()

}
