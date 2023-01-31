package main

import (
	"errors"
	"fmt"
	// "io/ioutil"
)

func main() {
	// content, err := ioutil.ReadFile("./hello.txt")
	// if err != nil {
	// 	fmt.Printf("Ocurrio un error: %v", err)
	// 	return
	// } 

	// 	fmt.Println(string(content))
	result, err := division(10, 2)
	if err != nil {
		fmt.Printf("Ocurrio un error: %v", err)
		return
	} 

		fmt.Println(result)
}

// Esta segunda forma de hacerlo, es un poco tedioso y no recomendable ya que cuando las funciones son un poco grandes , es mas tedioso y 
// dificil de leer.

// Otra funcionalidad que tiene Go , es que los valores que retornamos en las funciones, se les
// puede asignar un nombre, se llaman parametros nombrados
func division(dividendo, divisor int) (result int, err error) {
	// if divisor == 0 {
	// 	return 0, errors.New("No puedes dividir por cero")
	// }

	//  No hace falta asignarle algun valor, ya que con los eh asignado como variable,
	// result retorna por defecto su valor 0

	if divisor == 0 {
		err = errors.New("No puedes dividir por cero")
		return 
	}

	result = dividendo / divisor
	return 
}