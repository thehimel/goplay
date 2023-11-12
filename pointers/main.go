package main

import "fmt"

func main() {
	num := 100
	var pointerInt *int

	pointerInt = &num
	fmt.Println("&num =", &num)
	fmt.Println("pointerInt =", pointerInt)

	fmt.Println("num =", num)
	fmt.Println("*pointerInt =", *pointerInt)

	fmt.Println("&pointerInt =", &pointerInt)
}
