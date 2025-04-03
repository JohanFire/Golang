/*
Comparison operators
== Equal to
!= Not equal to
> Greater than
< Less than
>= Greater than or equal to
<= Less than or equal to
*/

package main

import "fmt"

func main() {

	var x int = 10

	value := x < 6

	fmt.Println(value) // false
	// or
	fmt.Printf("%t \n", x < 6) // false

	fmt.Println()

	var num1 int = 10
	var num2 int = 20

	if num1 == num2 {
		fmt.Println("num1 is equal to num2")
	} else if num1 != num2 {
		fmt.Println("num1 is not equal to num2")
	} else if num1 > num2 {
		fmt.Println("num1 is greater than num2")
	} else if num1 < num2 {
		fmt.Println("num1 is less than num2")
	} else if num1 >= num2 {
		fmt.Println("num1 is greater than or equal to num2")
	} else if num1 <= num2 {
		fmt.Println("num1 is less than or equal to num2")
	} else {
		fmt.Println("else")
	}
}
