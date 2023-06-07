package main

func main() {
	// all vars have a initial value, there are no null or undefined

	// var variable_name "data_type"
	var x, y, z int // int's initial value is "0"
	var name string // strings initial value is an empty string
	var active bool // bool's initial value is false
	var cadena []string

	// var num int
	// num = 23
	// you can resume, declare & assign value to a variable in the same line
	// with this syntax, there is no neccessity to asign the data type, the compiler assigns it in real time
	num := 23
	num = 20 // give a new value to an already existing variable

	// when we declare a variable, Go is waiting for it to be used, you can not leave it there without using it.
	// so let's print them
	println("int x:", x,
		"\nint y:", y,
		"\nint z:", z,
		"\nint num:", num,
		"\nstring name:", name,
		"\nbool active:", active,
		"\n[]string cadena:", cadena) // so let's print it to not leave them without any use
}
