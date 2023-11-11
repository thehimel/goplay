package main

import (
	"fmt"
	"math"
)

func main() {
	left := 12
	right := 12.2345
	sum := float64(left) + right
	fmt.Println("Sum =", sum)
	fmt.Println("Sum (Till 2 decimal points) =", math.Round(sum*100)/100)
	fmt.Printf("Sum (Till 2 decimal points) = %.2f", sum)
	fmt.Printf("Value of √2 = %v, Type = %T", math.Sqrt2, math.Sqrt2)
}
