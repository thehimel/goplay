package main

import "fmt"

func main() {
	// Create a slice
	mySlice := []string{"Vanilla", "Chocolate", "Strawberry"}
	fmt.Println(mySlice)

	// Append an element to the slice
	mySlice = append(mySlice, "Butterscotch", "Caramel")
	fmt.Println(mySlice)

	// Print elements
	fmt.Println(mySlice[:])
	fmt.Println(mySlice[2:])
	fmt.Println(mySlice[2 : len(mySlice)-1])
	fmt.Println(mySlice[:len(mySlice)-1])

	// Make a slice
	fruits := make([]int, 4, 6)
	fmt.Println(fruits)

	// Append elements
	fruits = append(fruits, 2, 4, 6, 7)
	fmt.Println(fruits)
}
