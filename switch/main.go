package main

import "fmt"

func main() {
	// no hay necesidad de usar break por cada condicion
	// go lo hace automaticamente rompiendo el flujo
	// Cuando usemos operadores logicops y de comparacion en la sentencia switch
	// no debemos colocar la expresion que queremos evaluar
	emoji := "🥞"
	switch {
	case emoji == "🐱"|| emoji =="🐗":
		fmt.Println("Es un animal")
	case emoji =="🍨" || emoji == "🌮" :
		fmt.Println("Es una fruta")
	default:
		fmt.Println("No es un animal ni una fruta")
	}
}