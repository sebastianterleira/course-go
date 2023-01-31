package main

import "fmt"

// Funciones anonimas autoejecutables y por variable tipo arrow fuction en JS
func main() {
	// func() {
	// 	fmt.Println("Hello World")
	// }()
	x := func(text string) {
		fmt.Println("Hello " + text)
	}

	x("Sebastián")
}