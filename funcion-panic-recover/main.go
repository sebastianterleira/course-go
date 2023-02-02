package main

import "fmt"

func main() {
	division(2, 6)
	division(2, 12)
	division(20, 0)
	division(40, 10)
}

func division(dividendo, divisor int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recuperandome del panic:", r)
		}
	}()
	vilidarDivisor(divisor)
	fmt.Println(dividendo / divisor)
}

func vilidarDivisor(divisor int) {
	if divisor == 0 {
		panic("🎇")
	}  
}