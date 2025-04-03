package main

import "fmt"

func main() {
	var number int64 = 10

	switch number {
	case 1: // if number == 1
		fmt.Println("Number is 1")
	case 2, -2: // if number == 2 or number == -2
		fmt.Println("Number is 2 or -2")
	case 5:
		fmt.Println("Number is 5")
	// This will not work because number is int64 and "hello" is string
	// case "hello":
	default: // if none of the above cases match
		fmt.Println("Number is not 1, 2, -2 or 10")
	}

	fmt.Println()

	switch {
	case number < 5:
		fmt.Printf("Number: %d is less than 5 \n", number)
	case number > 5:
		fmt.Printf("Number: %d is greater than 5 \n", number)
	case number == 5:
		fmt.Println("Number is equal to 5")
	default:
		fmt.Println("Number is not less than, greater than or equal to 5")
	}
}
