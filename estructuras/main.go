package main

import "fmt"

// las estructuras en go me permiten almacenar
// una coleccion de campos, similar a las clases

func main() {
	type course struct {
		Name      string
		Professor string
		Country   string
	}

	db := course{
		Name: "Bases de datos",
		Professor: "Alexys",
		Country: "Colombia",
	}

	// Estructura Literal

	// git := course{"Git", "Beto", "Bolivia"}

	// Si quiero solo agregar un valor a solo un campo, 
	// uso el primer metodo, pero el resto de campos quedaran en valor 0
	// css := course{Professor: "Alvaro"}

	// Tambien se puede ingresar a un campo , y mostrar su valor
	// p := &db
	// (*p).Professor = "Beto"
	// Cambiar el valor , de un campo por punteros
	p := &db
	p.Professor = "Beto"

	// para casos con estructuras ,no hace falta hacer desreferenciacion
	// porque go lo hace automaticamente

	fmt.Printf("%+v\n\n", db)
	fmt.Printf("%+v\n\n", p)

	// fmt.Printf("%+v\n\n", git.Country)
	// fmt.Printf("%+v\n\n", css.Professor)
}
