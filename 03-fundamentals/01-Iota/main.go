package main

import "fmt"

func main() {

	// CONST can be declared inside or outside the function body
	const (
		a = 23
		b = 99
		c = iota
		d
		e = iota
	)

	// new CONST with iota, starts with 0(ZERO)
	const (
		w = iota
		x
		y
		z
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	fmt.Println("")

	fmt.Println(w)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
