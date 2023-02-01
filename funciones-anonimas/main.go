package main

import "fmt"

// Funciones anonimas autoejecutables y por variable tipo arrow fuction en JS
func main() {
	// func(text string) {
	// 	fmt.Println("Hello " + text)
	// }("Sebastián")
	x := func(text string) {
		fmt.Println("Hello " + text)
	}

	x("Sebastián")
}