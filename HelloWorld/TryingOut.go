package main

import "fmt"

func main() {
	fmt.Println("Try out something !")

	foo()

	for i := 0; i < 5; i++ {
		fmt.Print("Iteration ", i)

		if i%2 == 0 {
			fmt.Print(" - Even iteration !")
		} else {
			fmt.Print(" - Odd iteration !")
		}

		fmt.Println()
	}
}

func foo() {
	fmt.Println("We are in foo")
}
