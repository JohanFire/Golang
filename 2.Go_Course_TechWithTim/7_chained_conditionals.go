/*
Chained Conditionals:
==
!=
&&
||

*/

package main

import "fmt"

func main() {
	if 5 > 3 && 3 > 1 {
		fmt.Println("5 is greater than 3 and 3 is greater than 1")
	}

	if 5 > 3 || 3 > 1 {
		fmt.Println("5 is greater than 3 or 3 is greater than 1")
	}

	if !true == false {
		fmt.Println("!true == false")
	}

	if !!false == false {
		fmt.Println("!!false == false")
	}
}
