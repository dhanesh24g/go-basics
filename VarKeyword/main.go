package main

import "fmt"

// var is getting used for DECLARATION as well as ASSIGNMENT
// var keyword outside function body is GLOBAL
var z = "Z - Getting DECLARED & ASSIGNED 'outside' the Function body.."

// declare a variable with THE ZERO VALUE
var a int
var b string
var c bool

func main() {

	// short declaration
	x := 22
	fmt.Println("X - ", x)

	// var keyword inside function body
	var y = "Y - Getting DECLARED & ASSIGNED 'inside' the function body.."
	fmt.Println(y)

	fmt.Println(z)

	// only ASSIGN (decared above)
	z = "Z - Getting ASSIGNED 'inside' the function body.."
	fmt.Println(z)

	foo()

	// testing ZERO VALUE for variables
	fmt.Println("Testing ZERO VALUE for variables - ")
	fmt.Println("A - ", a)
	fmt.Println("B - ", b)
	fmt.Println("C - ", c)
}

func foo() {

	z = "Z - Getting assigned in function FOO.."
	fmt.Println(z)
	fmt.Println()
}
