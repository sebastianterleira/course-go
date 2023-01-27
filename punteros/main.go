package main

import "fmt"

func main() {
	fruit := "🍎"
	var p *string
	p = &fruit
	*p = "🍓"
	fmt.Printf("Tipo: %T, Valor: %v, Direccion: %v\n", fruit, fruit, &fruit)
	fmt.Printf("Tipo: %T, Valor: %v, Desreferenciacion: %s\n", p, p, *p)
}