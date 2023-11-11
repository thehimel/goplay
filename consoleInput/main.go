package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter input: ")
	data, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading from console!")
		return
	}
	fmt.Println("Entered data:", data)
}
