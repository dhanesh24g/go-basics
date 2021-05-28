package main

import (
	"fmt"
)

func main() {
	x := 0

	// there is no WHILE loop in GO
	// FOR loop written like WHILE
	fmt.Println("Loop 1 -")
	for x < 4 {
		fmt.Println(x)
		x++
	}

	// FOR loop with BREAK statement
	fmt.Println("\nLoop 2 -")
	for {
		if x < 0 {
			break
		}
		fmt.Println(x)
		x--
	}

	// Printing even numbers from -1 to 100
	fmt.Println("\nPrinting EVEN numbers -")
	for {
		x++
		if x > 100 {
			break
		}

		if x%2 != 0 {
			continue
		}
		fmt.Printf("%v ", x)
	}
}
