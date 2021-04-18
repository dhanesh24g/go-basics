package main

import "fmt"

var i = 22

// back quote is a RAW STRING LITERAL
// ` would take everything as a string in between
var s = `I said,
	      "I achieve all my goals !"`

// defining boolean type variable
var b = false

func main() {

	fmt.Print("i with Value - ", i, " & Type - ")
	fmt.Printf("%T\n", i)

	fmt.Println("s with Value - ", s)

	fmt.Print("b with Value - ", b, " & Type - ")
	fmt.Printf("%T\n", b)
}
