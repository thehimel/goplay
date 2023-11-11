package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Now =", time.Now())
	createdMoment := time.Date(2023, time.December, 10, 14, 44, 22, 8, time.UTC)
	fmt.Println("Custom Moment =", createdMoment)
	fmt.Println("Custom Moment =", createdMoment.Format(time.DateTime))
}
