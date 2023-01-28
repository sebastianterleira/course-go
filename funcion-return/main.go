package main

import "fmt"

func main() {
	total := sum(2, 3)
	fmt.Println("Return:", total)
}

func sum(a, b int) int{
	return a + b
}