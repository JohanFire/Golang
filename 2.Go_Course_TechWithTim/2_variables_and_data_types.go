package main

import "fmt"

func main() {
	// all vars have a initial value, there are no null or undefined
	// there are Implicit and Explicit assignment expressions

	var name string // strings initial value is an empty string
	var num1 int    // int's initial value is "0"
	var num2, num3 int
	var num4 int16 = 23 // this is an explicit assignment expression

	// you can resume, declare & assign value to a variable in the same line
	// with this syntax, there is no neccessity to asign the data type, the compiler assigns it in real time
	num5 := 23 // implicit assignment expression

	var isAnAdult bool // bool's initial value is false

	fmt.Println("Enter your name:")
	fmt.Scanln(&name)

	fmt.Println("Hi,", name, "!", "\nAre you an adult?:", isAnAdult)
	fmt.Println()

	fmt.Println("My first number is:", num1)
	fmt.Println("My second number and third number are:", num2, num3)
	fmt.Printf("My second number and third number are: %d %d \n", num2, num3)
	fmt.Println("My fourth number is:", num4, "and its of type: %T", num4)
	fmt.Println("My fifth number is:", num5)
	fmt.Printf("My fifth number is of type: %T\n", num5)

	fmt.Println()

	fmt.Println("My fourth number is:", num4)
}
