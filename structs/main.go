package main

import "fmt"

type Fruit struct {
	name   string
	weight float64
}

func main() {
	apple := Fruit{"Apple", 100}
	fmt.Println(apple)
	fmt.Println(apple.name)

	apple.name = "Green Apple"
	fmt.Println(apple.name)
}
