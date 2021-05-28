package main

import "fmt"

func main() {

	var name string
	fmt.Printf("\nInput your name: ")
	fmt.Scanln(&name)

	switch name {
	case "Dhanesh":
		fmt.Println("Hello Dhanesh, you are awesome !")
		//	fallthrough
	case "Parveen":
		fmt.Println("Hello Parveen, you are awesome !")
	default:
		fmt.Println("Sorry, don't know you but still you are awesome !")
	}
}
