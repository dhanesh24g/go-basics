package main

import "fmt"

func printValue(a int) {

	fmt.Print("\nValue of the variable in -\n\nDecimal\t\tBinary\t\tHexaDecimal\n")
	fmt.Printf("%d\t\t%b\t\t%#X\n", a, a, a)
}

func main() {

	fmt.Print("\nProvide Input: ")
	var y int
	// Take input from user
	fmt.Scanln(&y)

	printValue(y)

	fmt.Printf("\n")

	// using RAW STRING LITERAL ``
	s := `	<<<
	Shifting the bits by 1 position to left
	<<<`

	fmt.Print(s, "\n")

	// Declaring TYPED & UNTYPED const
	const (
		a     = 42
		b int = 22
	)

	y = y << 1

	printValue(y)
}
