package main

import "fmt"

func main() {
	var sum int
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Printf("Count sum: %d\n", sum)
}
