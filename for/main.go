package main

import "fmt"

// En go solo existe una instruccion de ciclo , que es for pero 
// este tiene diferentes variantes

func main() {
	// count := 10
	// for i := 1; i <= count; i++ {
	// 	fmt.Println(i)
	// }


	// Ciclo for Continuo es similar al ciclo while en otros lenguajes
	// lo que me dice go es que la declaracion de inicio, como la declaracion 
	// posterior son opcionales

	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++ 
	// }

	// Ciclo forever como su nombre indica , es un ciclo que dura para siempre
	// se usa para  procesos que necesitan ser escuchados permanentemente, como sockets
	
	// i := 1
	// for {
	// 	fmt.Println(i)
	// 	i++ 
	// }

	// Ciclo for-range-slice permite iterar sobre slice, maps, y tambien strings

	// nums := []uint8{1, 2, 3, 4}

	// for i := range nums {
		// fmt.Println("Indice:", i, "Valor:", v)
	// 	nums[i] *= 2
	// }
	// fmt.Println(nums)

	// Ciclo For-range-map

	// sports := map[string]string{"Basketball": "🏀", "soccer": "⚽"}

	// for key, v := range sports {
	// 	fmt.Println("key:", key, "Valor:", v)
	// }

	// Ciclo For-range-string

	hello := "hello"

	for i, v := range hello {
		fmt.Println("Indice:", i, "Valor:", string(v))
	} 

}