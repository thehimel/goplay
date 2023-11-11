package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a number: ")
	data, errData := reader.ReadString('\n')
	if errData != nil {
		log.Fatalln("Error reading input from console.")
	}

	num, errNum := strconv.ParseFloat(strings.TrimSpace(data), 64)
	if errNum != nil {
		log.Fatalln("Error converting data to number for: ", errNum)
	}
	fmt.Println("Entered number: ", num)
}
