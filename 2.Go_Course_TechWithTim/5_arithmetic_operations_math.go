/*
Order of operations in math:
PEMDAS
P: Parentheses
E: Exponents
M: Multiplication
D: Division
A: Addition
S: Subtraction

In Go, the order of operations is the same as in math.
The order of operations is important to remember when writing code.
If you don't follow the order of operations, you may get unexpected results.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 int16 = 10
	var num2 int16 = 3

	answer := num1 + num2
	fmt.Printf("The sum of %d and %d is: %d \n", num1, num2, answer)
	fmt.Println()

	var num3 float64 = 3.0
	var num4 int16 = 20

	answer2 := num3 * float64(num4)

	fmt.Printf("The product of %f and %d is: %.2f \n", num3, num4, answer2)
	fmt.Println()

	answer3 := float64(num4) / num3
	fmt.Printf("The division of %d and %f is: %.2f \n", num4, num3, answer3)
	fmt.Println()
	fmt.Println()

	// Using the math package to calculate the square root of a number
	// number := 16.0
	var number float64 = 16.0
	sqrtResult := math.Sqrt(number)
	fmt.Printf("The square root of %.2f is: %.0f \n", number, sqrtResult)
	fmt.Println()

	// Using the math package to calculate the power of a number
	var base float64 = 2
	exponent := 3.0
	powerResult := math.Pow(base, exponent)
	fmt.Printf("%.2f raised to the power of %.2f is: %.2f \n", base, exponent, powerResult)
	fmt.Println()
}
