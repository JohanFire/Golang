package main

import "fmt"

func main() {
	// for loop
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}

		fmt.Println("Value of i:", i)
	}

	fmt.Println()

	// while loop (using for loop syntax)
	i := 0
	for i < 5 {
		fmt.Println("Value of i:", i)

		i++
	}

	fmt.Println()

	// infinite loop (not recommended, but possible)
	// for true{
	for { // same as for true
		fmt.Println("Infinite loop = equivalent to while true")
		break // to prevent an actual infinite loop in this example

	}

}
