package main

import "fmt"

func main() {
	// defer fmt.Println(3)
	// defer fmt.Println(2)
	// defer fmt.Println(1)

	// Particulariadad con Defer: los argumentos de la funcion son evaluados inmediatamente,
	// pero se aplaza la ejecucion, es por eso que el valor de a es 5 porque es el valor que se tenia en el momento y para la segunda impresion al cambiarlo a 10 te muestra 10
	a := 5 
	defer fmt.Println("Valor:",a)

	a = 10
	fmt.Println(a)
}