package main

import "fmt"

func main() {

	// Declare slice
	slc := []int{21, 3, 4, 64, 2}

	fmt.Printf("%v \n", slc)
	fmt.Printf("Slice Length : %v \n\n", len(slc))

	// Range of slice
	fmt.Println("Range of slice -")
	for i, v := range slc {
		fmt.Println(i, v)
	}

	// Slice of a slice
	fmt.Printf("\nSlice of slice -\n")
	fmt.Println(slc[:])   // Complete Slice
	fmt.Println(slc[1:4]) // Start from 1 & go upto 4 (Don't include 4)

	// Make Slice on top of Array
	// make([]T, length, capacity)
	newSlc := make([]int, 5, 10)

	// newSlc[5] = 21 will give Index out of Range
	newSlc = append(newSlc, 23, 43)
	fmt.Printf("\nSlice: %v\n", newSlc)
	fmt.Println("Length of Slice: ", len(newSlc))
	fmt.Println("Capacity of Array: ", cap(newSlc))

	// If we keep appending the elements & the capacity is crossed
	// Complete Array is shifted to new place & that's by internal computing
	newSlc = append(newSlc, 53, 63, 73, 83, 93)
	fmt.Println("\nNew Slice: ", newSlc)
	fmt.Println("New length of Slice: ", len(newSlc))
	fmt.Println("New capacity of Array: ", cap(newSlc))
	fmt.Println("Since we crossed the original Capacity, the new Capacity is doubled with internal computation !")
}
