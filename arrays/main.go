package main

import "fmt"

func main() {
	var fruits [3]string
	fruits[0], fruits[1], fruits[2] = "Apple", "Orange", "Banana"

	fmt.Println(fruits)
	for i := 0; i < len(fruits); i++ {
		delimiter := ""
		if i == len(fruits)-1 {
			delimiter = "\n"
		}
		fmt.Printf("%v %v", fruits[i], delimiter)
	}

	var digits = [4]int{1, 2, 3, 4}
	fmt.Println(digits)
}
