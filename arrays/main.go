package main

import "fmt"

func main() {

	/*
	var food [3]string
	food[0] = "🍿"
	food[1] = "🍔"
	food[2] = "🍕"
	*/
	// Error, los array no se puede ampliar,
	// el tamaño de un array es fijo
	// food[3] = "🍕"

	// ARRAYS LITERALES

	food := [3]string{"🍿", "🍔",  "🍕"}

	fmt.Printf("Tipo: %T, Valor: %v", food, food)
}