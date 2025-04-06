/*
An interface is a type that defines a set of method signatures.
An interface specifies a behavior that a type must implement,
but it does not provide the implementation of those methods.

This allows different types to satisfy
the same interface as long as they implement the required methods.
*/

package main

import (
	"fmt"
	"math"
)

// Define the interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Type that implements the interface: Circle
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Type that implements the interface: Rectangle
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Function that takes the interface as a parameter
func get_area(s Shape) float64 {
	return s.Area()
}

func main() {
	// Create instances of Circle and Rectangle
	circle1 := Circle{Radius: 5}
	rectangle1 := Rectangle{Width: 4, Height: 6}

	// Use the Shape interface
	fmt.Println("Using the Shape interface:")
	var shapeInterface Shape

	shapeInterface = circle1
	fmt.Printf("Circle Area: %.2f, Perimeter: %.2f\n", shapeInterface.Area(), shapeInterface.Perimeter())

	shapeInterface = rectangle1
	fmt.Printf("Rectangle Area: %.2f, Perimeter: %.2f\n", shapeInterface.Area(), shapeInterface.Perimeter())

	fmt.Println()

	// Create a slice of Shape interface
	fmt.Println("Using a slice of Shape interface:")

	shapes := []Shape{circle1, rectangle1}

	fmt.Println(shapes)
	fmt.Println("Circle Area:", shapes[0].Area())
	for _, shape := range shapes {
		fmt.Printf("Shape Area: %.2f, Perimeter: %.2f\n", shape.Area(), shape.Perimeter())
	}

	fmt.Println()

	// Use the interface as a parameter
	fmt.Println("Using the interface as a parameter:")

	for _, shape := range shapes {
		fmt.Println("Area:", get_area(shape))
	}

	fmt.Println()
}
