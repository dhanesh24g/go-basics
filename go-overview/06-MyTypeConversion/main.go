package main

import "fmt"

type mytype int

var x mytype

func main() {

	x = 32
	fmt.Println("X = ", x)
	fmt.Printf("Type of x is - %T\n", x)

	y := 1033
	// x = y
	// will not work, as both have different TYPES
	fmt.Println("Y = ", y)

	// type can be converted like T(x)
	y = int(x)
	fmt.Println("Y (after Conversion) = ", y)

	// type can be converted like T(x)
	y = 2099
	x = mytype(y)
	fmt.Println("X (after Conversion) = ", x)
}
