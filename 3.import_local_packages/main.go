package main

import (
	"fmt"

	"johanfire.com/imports/third" // Importing the third package
)

func main() {
	fmt.Println("Hello World", Name)
	// Name is a public variable in the second.go file
	// can be accessed without importing the package
	// because it is in the same package "main"
	// means it is in the same directory
	// so if you run "go run ." it will work

	fmt.Println()

	// ######################################################
	// Now importing packages from another directory
	// you must instance always the package name,
	// so it could be posible that the 'package' name
	// inside the .go file is different from the directory name
	// e.g
	// "johanfire.com/imports/third" is the directory name
	// the directory name is "/third"
	// and the package name is "myThirdPackage"
	// so you must instance the package name, "myThirdPackage"
	// fmt.Println(myThirdPackage.MyThirdPackageName)
	// fmt.Println(myThirdPackage.myThirdPackageLastName) // Private variable from third package

	// Although the optimal, is to use the same name for the package and the directory
	fmt.Println(third.MyThirdPackageName)
}
