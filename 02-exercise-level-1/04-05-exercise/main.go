package main

import "fmt"

func main() {

	// my own type with UNDERLYING type as int
	// scope of mytype is only in this function
	type mytype int
	var x mytype

	fmt.Printf("Initial value of x: %v\nType of variable x: %T\n", x, x)

	x = 42
	// traditional way
	// fmt.Println("New value of x:", x)
	fmt.Printf("New value of x: %v\n\n", x)

	// my second own type with UNDERLYING type as mytype
	type newtype mytype
	var z newtype

	fmt.Printf("Initial value of z: %v\nType of variable z: %T\n\n", z, z)
	foo()

	// SHORT DECLARATION with type CONVERSION
	y1 := int(x)
	fmt.Printf("Initial value of y1: %v\nType of variable y1: %T\n\n", y1, y1)

	// SHORT DECLARATION without type CONVERSION
	y2 := x
	fmt.Printf("Initial value of y2: %v\nType of variable y2: %T\n\n", y2, y2)

	// type CONVERSION
	z = newtype(x)
	fmt.Printf("New value of z: %v\nType of variable z: %T\n\n", z, z)
}

func foo() {

	// type footype mytype
	// above statement will not compile, because mytype is not in scope of foo
	type footype int
	var a footype

	fmt.Printf("Initial value of a: %v\nType of variable a: %T\n\n", a, a)
}
