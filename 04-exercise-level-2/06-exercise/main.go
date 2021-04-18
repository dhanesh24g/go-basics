package main

import "fmt"

func main() {

	const (
		a = 2021 + iota
		b // (2021+iota) IOTA is now 1
		c // (2021+iota) IOTA is now 2
		d // (2021+iota) IOTA is now 3
		//IOTA's next value will be 4
		e = 2025 - iota
		f // (2025-iota) IOTA is now 5
		g // (2025-iota) IOTA is now 6
		h // (2025-iota) IOTA is now 7
	)
	fmt.Println("Current & next three years -")
	fmt.Print(a, b, c, d, "\n\n")

	fmt.Println("Current & last three years -")
	fmt.Println(e, f, g, h)
}
