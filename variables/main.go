package main

import "fmt"

const fruitName string = "Apple"

func main() {
	fmt.Println(fruitName)

	digit := 100
	fmt.Println(digit)

	var word = "Simple"
	fmt.Println(word)
	fmt.Printf("Variable type: %T.\n", word)

	var num int64 = 10000000000
	fmt.Printf("Value = %v, Type = %T", num, num)
}
