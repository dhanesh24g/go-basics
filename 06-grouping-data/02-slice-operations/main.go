package main

import "fmt"

func main() {

	slc := []int{21, 45, 32, 4, 99, 12}

	// Append to Slice
	slc = append(slc, 100, 123, 432)
	fmt.Printf("Appended to Slice: \n%v\n\n", slc)

	slcNew := []int{1, 3, 5, 7}
	fmt.Printf("New Slice: \n%v\n\n", slcNew)

	slcNew = append(slcNew, slc...)
	fmt.Printf("Appended New Slice to Old Slice: \n%v\n\n", slcNew)

	// Delete from Slice
	slcDel := append(slc[:3], slc[5:8]...)
	fmt.Printf("Deleted from Old Slice: \n%v\n\n", slcDel)

	// Multi-dimensional Slice
	name := []string{"Dhanesh", "Saumitra", "Ankush", "Pratik"}
	surname := []string{"Gujrathi", "Kahate", "Shingane", "Singh"}
	fullname := [][]string{name, surname}

	fmt.Printf("Fullname Multi-Dimensional Slice: \n%v", fullname)
}
