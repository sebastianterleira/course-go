package main

import "fmt"

// En go solo existe una instruccion de ciclo , que es for pero 
// este tiene diferentes variantes

func main() {
	count := 10
	for i := 1; i <= count; i++ {
		fmt.Println(i)
	}

}