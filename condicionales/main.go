package main

import "fmt"

func main() {
/*	emoji := "👣"
	if emoji == "👣" {
		fmt.Println("es un cactus")
	} else if emoji == "😋" {
		fmt.Println("es una carita")
	} else {
		fmt.Printf("el emoji es: %s", emoji)
	} */
	
	// que particularidad tiene esto ?, que cuando se declara la variable dentro de la
	// estructura de control if, la variable se puede utilizar solo dentro de esa estructura if
	// mas no fuera de ella.
	if emoji := "👣"; emoji == "👣" {
		fmt.Println("es un cactus")
	} else if emoji == "😋" {
		fmt.Println("es una carita")
	} else {
		fmt.Printf("el emoji es: %s", emoji)
	}

	// Error ya que la variable emoji solo funciona dentro del scope de la 
	// estructura de control if
	// fmt.Printf(emoji)
}