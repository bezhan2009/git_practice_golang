package main

import "fmt"

func sum(a uint, b uint) uint {
	return a + b
}

func main() {
	fmt.Println("Sum:", sum(1, 2))
}
