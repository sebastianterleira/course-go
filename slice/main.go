package main

import "fmt"

// los slice, nos ayudan a trabajar con arreglos de forma dinamica

func main() {
	// Slice no poseen datos, son apuntadores a un Array 
	set := [7]string {"🍟", "🌭", "🥗", "🧂", "🥐", "🥓", "🥚"}
	// Tambien si no se pasa el inicio go empezara desde el indice 0, igualmente con final
	// y si no especificamos ningun, pues se muestra todo el array.
	animals := set[:4]
	fly := set[5:7]
	fly[0] = "🍭"
	fmt.Printf("Array => Tipo: %T, Valor: %v\n\n", set, set)
	fmt.Printf("Sin especificar el inicio => Tipo: %T, Valor: %v\n\n", animals, animals)
	fmt.Printf("Tipo: %T, Valor: %v\n\n", fly[0], fly[0])
	fmt.Printf("All => Tipo: %T, Valor: %v\n\n", fly[0], fly[:])
}