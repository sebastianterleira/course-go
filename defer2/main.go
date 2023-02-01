package main

import (
	"fmt"
	"os"
)

// Usualmente Defer se usa para limpiar recursos,
// para cerrar archivos, para cerrar controladores de red o cerrar controladores de Base de Datos

func main() {
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Printf("Ocurrio un error al crear el archivo: %v", err)
		return
	}
	defer file.Close()

	_, err = file.Write([]byte("Hello Codeablers"))
	if err != nil {
		fmt.Printf("Ocurrio un error al escribir el archivo: %v", err)
		return
	}
}
