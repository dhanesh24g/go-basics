package main

import (
	"fmt"
)

func main() {

	fmt.Println("Decimal Binary")
	x := 4
	fmt.Printf("%v\t%b\n", x, x)

	// << & >> are bit shifting operators
	y := x << 1
	fmt.Printf("%d\t%b\n", y, y)

	z := y >> 1
	fmt.Printf("%d\t%b\n\n", z, z)

	// IOTA for bit shifting
	const (
		_  = iota
		kb = 1 << (10 * iota)
		//mb = 1 << (10 * iota)
		//gb = 1 << (10 * iota)
		mb
		gb
	)

	fmt.Println("Decimal \tBinary")
	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)

}
