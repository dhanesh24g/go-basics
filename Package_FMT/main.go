package main

import "fmt"

func main() {

	x := 15
	fmt.Println("Normal Println statement..")

	// values can be formatted in various Numerical Systems with Printf
	fmt.Printf("Boolean value of %v - %b \n", x, x)
	fmt.Printf("Hex value of %v - %x \n", x, x)
	fmt.Printf("Hex representation of %v - %#x \n\n", x, x)

	// string value will be assigned to s like below
	s := fmt.Sprint("Value will be assigned from here, with Sprint..")
	fmt.Println(s)

	y := 22
	num := fmt.Sprintf("Boolean value of %v - %b \n", y, y)
	fmt.Println(num)
}
