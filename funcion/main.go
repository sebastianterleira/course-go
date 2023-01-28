package main

import "fmt"

func main() {
	// hello("Sebastian", "Terleira")
	emoji := "🎁"
	change(&emoji)
	fmt.Printf(emoji)
}

// func hello() {
// 	fmt.Println("Hello world")
// }

// func hello(firstName string, lastName string) {
// 	fmt.Printf("Hello %s %s",firstName, lastName)
// }

func change(value *string) {
	*value = "🎢"
}
