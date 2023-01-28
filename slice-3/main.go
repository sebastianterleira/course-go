package main

import "fmt"

func main() {

	// Arrays Literales
	// cuando no se le pasa ningun tamaño fijo, go lo interpreta con un slice
	// sin ninguna referencia hacia un arreglo, asi que slice lo primero que hace es
	//  crear su propio arreglo en donde arreglo los dos elementos que ya existen.
	// append retorna un slice que referencia a ese array que el creo.
	
	// fruits := []string{"🍟", "🥗"}

	// Make
	// construccion de arrays con la funcion preconstruida make()

	// fruits := make([]string, 0, 3)
	
	// fruits = append(fruits,"🍟", "🥗", "🥤", "🍎")

	// fmt.Printf("FRUITS => Tipo: %T, Valor: %v\n\n", fruits, fruits)
	// fmt.Printf("FRUITS => Tipo: %T, Valor: %v\n\n", fruits, len(fruits))
	// fmt.Printf("FRUITS => Tipo: %T, Valor: %v\n\n", fruits, cap(fruits))

	// Valor 0 de los slices

	// automaticamente el valor cera 0, el valor 0 de las frutas es nil
	// en otros lenguajes "nulo"
	
	// var fruits []string

	// pero si hacemos uno literal aunque no posea elementos, recibiriamos falso , porque ya no seria nulo
	
	var fruits = []string{}
	fmt.Printf("FRUITS => Tipo: %T, Valor: %v\n\n", fruits, fruits)
	fmt.Printf("FRUITS => Tipo: %T, Valor: %v\n\n", fruits, len(fruits))
	fmt.Printf("FRUITS => Tipo: %T, Valor: %v\n\n", fruits, cap(fruits))

	// Verifiquemos si es nil
	fmt.Println(fruits == nil)
}