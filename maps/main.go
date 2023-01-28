package main

import "fmt"

func main() {

	animals := make(map[string]string)
	animals["cat"] = "🐈"
	animals["dog"] = "🐕"

	fmt.Printf("Tipo: %T, Valor: %v\n\n", animals, animals)

	// Mapa literal 
	fruits := map[string]string{
		"apple": "🍎",
		"banana": "🍌",
	}
	fmt.Printf("Tipo: %T, Valor: %v\n\n", fruits, fruits)

	// Eliminar elementos con map, con la funcion preconstruida
	// delete()

	delete(fruits, "banana")
	fmt.Printf("Tipo: %T, Valor: %v\n\n", fruits, fruits)

	// como obtener los elemento

	// cuando se especifica una clave que no existe, devuelve un valor 0, osea
	// un string vacio.


	//Podemos hacer una validacion y decirle que si no existe
	// cree esa clave "gorilla"

	if _, ok := animals["gorilla"]; !ok {
		animals["gorilla"] = "🙈"
	}

	fmt.Println("Valor: \n\n", animals)
	fmt.Println("Valor: \n\n", animals["cat"])
}