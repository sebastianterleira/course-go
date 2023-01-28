package main

import "fmt"

func main() {
	//  len(): # de elementos en el slice
	// 	cap(): # de elementos del array donde apunta el slice, a
	// parti del indice de donde se creó el slice

	food := [5]string {"🍟", "🥗", "🧀", "🥟", "🍙"}
	fruits := food[1:3] // "🥗", "🧀"

	// recordemos que el slice es dinamico, asi que se puede agragar
	// mas elementos segun la cantidad maxima que tenga el array

	// forma de hacerlo no valida(el compilador dara error)
	// fruits[2] = "🍚"

	// Forma correcta es con la funcion pre construida
	// de go que se llama append()

	// el primer argumento recibe el slice donde quiero agregar elementos
	// y con el segundo argumento, los n elementos que agregare
	fruits = append(fruits, "🍧", "🥕", "🥠")

	// cuando se sobrepasa la cantidad de elementos, ocurren peculiaridades
	// el arreglo original que es "FOOD", ya no cambia el valor y la cantidad
	// del arreglo de "FRUITS" pasa a ser 8... esto, por que?

	// Cuando se sobrepasa la cantidad maxima del slice al que esta haciendo referencia
	// la funcion "append", go crea un nuevo array y duplica la cantidad del primer arreglo
	// hacia el nuevo arreglo entonces seria 8 la cantidad maxima, y los elementos que se modificaban 
	// en el arreglo original vuelve a la normalidad.

	// Array[4] {"🥗", "🧀", "🥟", "🍙"} se agrega uno mas y supera su cantidad maxima
	// new Array[8] {"🥗", "🧀", "🍧", "🥕", "🥠"} y pasa esto 

	// Lo que pasa luego es que append retorna un slice que apunta al nuevo array es decir, que pierde la referencia que tenia con el array anterior
	// la variable FOOD y retorna una referencia al nuevo array y los valores en la variable FOOD ya no
	// van a cambiar y podemos agregar los nuevos elementos, y asi sucesivamente go estara haciendo este ejercicio 
	// para poder trabajar de manera dinamica con los arrays, cuando agregue mas elementos y llega a la capacidad de 8 , go hara lo mismo 
	//  y ahora el nuevo array sera de capacidad de 16 elementos

	fmt.Printf("FOOD => Tipo: %T, Valor: %v\n\n", food, food)
	fmt.Printf("FRUITS => Tipo: %T, Valor: %v\n\n", fruits, fruits)
	fmt.Printf("FRUITS => Tipo: %T, Valor: %v\n\n", fruits, len(fruits))
	// la capacidad que se obtiene es a partir de donde 
	// inicia el slice, osea desde la posición 1 respectivamente
	// asta el final.
	fmt.Printf("FRUITS => Tipo: %T, Valor: %v\n\n", fruits, cap(fruits))
}