package main

import "fmt"

func main() {

	/*
	var food [3]string
	food[0] = "πΏ"
	food[1] = "π"
	food[2] = "π"
	*/
	// Error, los array no se puede ampliar,
	// el tamaΓ±o de un array es fijo
	// food[3] = "π"

	// ARRAYS LITERALES

	food := [3]string{"πΏ", "π",  "π"}

	fmt.Printf("Tipo: %T, Valor: %v", food, food)
}