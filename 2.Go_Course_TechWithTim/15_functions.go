package main

import "fmt"

func tests() {
	fmt.Println("tests")
}

// this means, all three params are of type int
// func add(num1, num2, num3 int) int {

func add(num1 int, num2 int) int {
	return num1 + num2
}

func add_and_substract(num1, num2 int) (int, int) {
	// return num1 + num2, num1 - num2

	sum := num1 + num2
	sub := num1 - num2
	return sum, sub

}

func multiply_and_divide(num1, num2 int) (response1 int, response2 float64) {
	defer fmt.Println("defer function called") // defer functions are executed at the end of the function, before returning

	response1 = num1 * num2
	response2 = float64(num1) / float64(num2)

	fmt.Println("before return")

	return // no need to add values here, because we already defined them above
}

func main() {
	fmt.Println("Functions in Golang")

	tests()

	fmt.Println()

	var num1 int = 5
	var num2 int = 10

	sum1 := add(num1, num2)
	fmt.Printf("Sum of %d and %d is: %d \n", num1, num2, sum1)

	fmt.Println()

	summatory, substraction := add_and_substract(num1, num2)
	fmt.Printf("%d + %d = %d \n", num1, num2, summatory)
	fmt.Printf("%d - %d = %d \n", num1, num2, substraction)

	fmt.Println()

	mult, div := multiply_and_divide(num1, num2)
	fmt.Printf("%d * %d = %d \n", num1, num2, mult)
	fmt.Printf("%d / %d = %.2f \n", num1, num2, div)
}
