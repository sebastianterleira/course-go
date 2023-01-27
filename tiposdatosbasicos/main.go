package main

import "fmt"

func main() {
	//  bool, string, numeric
	var boolean bool = true;
	var text string = "Hola";
	var num uint8 = 100;
	var Byte byte = 200;
	var Rune rune = -100;
	// valores Unicode con rune
	var Unicode rune = 'a';
	var floatx32 float32 = 23.44;
	var floatx64 float64 = 22.55;

	fmt.Printf("Tipos de datos , presentaciones:\n\n")

	fmt.Printf("Tipo: %T, Valor: %v\n", boolean, boolean)
	fmt.Printf("Tipo: %T, Valor: %v\n", text, text)
	fmt.Printf("Tipo: %T, Valor: %v\n", num, num)
	fmt.Printf("Tipo: %T, Valor: %v\n", Byte, Byte)
	fmt.Printf("Tipo: %T, Valor: %v\n", Rune, Rune)
	fmt.Printf("Tipo: %T, Valor: %v\n", Unicode, Unicode)
	fmt.Printf("Tipo: %T, Valor: %v\n", floatx32, floatx32)
	fmt.Printf("Tipo: %T, Valor: %v\n\n", floatx64, floatx64)


	fmt.Printf("Como utilizar Castin:\n\n")
	// No puedo operar con tipos de datos diferentes
	// // error por ser de tipos diferentes
	var a uint8 = 100
	var b uint16 = 2300
	// c := a + b

// solucion
	c := uint16(a) + b
	// no modificamos, transformamos el tipo de dato de "a" 
	// a el mismo tipo de valor de "b", pero recuerda,
	// no modifica el tipo de dato original que era de uint8
	// ya que respetamos el tipado fuerte de Go
	
	fmt.Printf("Tipo: %T, Valor: %v\n", c, c);
	// aun sigue siendo de tipo de dato uint8
	fmt.Printf("Tipo: %T, Valor: %v\n\n", a, a);

	// si queremos mantener una variable aunque no se use
	// usamos el identificador blank "_" 
	_ = 234;
	var _ string = "test_text"

	
	fmt.Printf("Valores por defecto en Go, peculiaridades 💀\n\n");
	// Valores por defecto en Go, peculiaridades 💀
	var textCero string 
	var numCero uint 
	var floatCero float32 
	var boolDefault bool 

	fmt.Printf("Valor: %q\n", textCero);
	fmt.Printf("Valor: %v\n", numCero);
	fmt.Printf("Valor: %v\n", floatCero);
	fmt.Printf("Valor: %v\n", boolDefault);
}