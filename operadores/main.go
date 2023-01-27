package main

import "fmt"

func main() {
	// Operadores Aritmeticos: (), *, /, %, +, -
	var a = 4 + 2*3

	fmt.Printf("Tipo: %T, Valor: %v\n\n", a, a)

	// Operadores de Asignacion: =, +=, -=, *=, /=, %=
	var b = 10
	b += 2

	fmt.Printf("Tipo: %T, Valor: %v\n\n", b, b)

	// declaracion post-incremeto y post-decremento: ++, --
	// (no son una expresion sino una declaracion)
	var c = 3
	c--

	fmt.Printf("Tipo: %T, Valor: %v\n\n", c, c)

	// Operadores de comparacion: >, <, ==, !=, >=, <=

	fmt.Println(5 != 3)

	// Operadores logicos: &&, ||
	var age = 60
	fmt.Println("niño", age < 18)
	fmt.Println("adulto:", age >= 18 && age <= 60)
	fmt.Println("viejo:", age > 60)

	//Unario: !
	fmt.Println(!(4 != 4))
	
}
