package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter input: ")
	inputData, errInput := reader.ReadString('\n')
	if errInput != nil {
		fmt.Println("Input error!")
		return
	}
	fmt.Println("Entered data:", inputData)
}
