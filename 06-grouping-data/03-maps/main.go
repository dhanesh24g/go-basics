package main

import "fmt"

var ch string
var orgMap map[string]string

func showChoice() {
	fmt.Printf("\nAvailable operations -\n1. Add Data\n2. Retrieve Data\n3. Remove Data\n4. Update Data\n5. Exit\n")
	fmt.Print("Pick the operation of your choice: ")
	var ch string
	fmt.Scanln(&ch)

	switch ch {

	case "1":
		fmt.Printf("\nAdding Record...\nEnter the person's name: ")
		addRecord()

	case "2":
		fmt.Printf("\nRetrieving Record...\nEnter the person's name: ")
		retrieveRecord()

	case "3":
		fmt.Printf("\nDeleting Record...\nEnter the person's name: ")
		delRecord()

	case "4":
		fmt.Printf("\nUpdating Record...\nEnter the person's name: ")
		updRecord()

	case "5":

	default:
		fmt.Printf("\nPlease enter a valid input !\nDo you want to exit ? (Y/N) : ")
		var e string
		fmt.Scanln(&e)
		if e == "n" || e == "N" {
			showChoice()
		}
	}
}

func main() {

	// Maps for key-value pair structure
	orgMap = map[string]string{
		"Dhanesh":  "Equinix",
		"Parveen":  "NXP",
		"Saumitra": "Deutsche Bank",
		"Ankush":   "Tibco",
	}

	//fmt.Println(orgMap)
	showChoice()
}

func addRecord() {

	var name, org string
	fmt.Scanln(&name)

	if _, ok := orgMap[name]; !ok {
		fmt.Printf("Enter the corresponding Organization name: ")
		fmt.Scanln(&org)
		orgMap[name] = org
		fmt.Printf("\nData added - %v works for %v\n", name, org)
	} else {
		fmt.Printf("\nName is already listed (Try updating the record if required) !\n")
	}
	showChoice()
}

func retrieveRecord() {

	var name string
	fmt.Scanln(&name)

	if _, ok := orgMap[name]; ok {
		fmt.Printf("\n%v works for %v\n", name, orgMap[name])
	} else {
		fmt.Printf("\nThis name is not listed !\n")
	}
	showChoice()
}

func delRecord() {

	var name string
	fmt.Scanln(&name)

	if _, ok := orgMap[name]; ok {
		delete(orgMap, name)
	} else {
		fmt.Printf("\nThis name is not listed !\n")
	}
	showChoice()
}

func updRecord() {

	var name, org string
	fmt.Scanln(&name)

	if _, ok := orgMap[name]; ok {
		fmt.Printf("Enter the corresponding Organization name: ")
		fmt.Scanln(&org)
		orgMap[name] = org
		fmt.Printf("\nData updated - %v works for %v\n", name, org)
	} else {
		fmt.Printf("\nThis name is not listed ! Please Add the record before updating !\n")
	}
	showChoice()
}
