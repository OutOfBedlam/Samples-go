package main

import (
	"fmt"

	"github.com/OutOfBedlam/Samples-go/mods/random"
)

func main() {
	r := random.New(10)

	var sum int
	for i := 1; i <= 10; i++ {
		sum += r.Next()
	}
	fmt.Printf("Count sum: %d\n", sum)
}
