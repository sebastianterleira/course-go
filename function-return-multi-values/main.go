package main

import (
	"fmt"
	"strings"
)

func main() {
	texto := "Sebastian"
	minusc, mayusc := convert(texto)
	fmt.Println(minusc, mayusc)
}

func convert(text string) (string, string) {
	min := strings.ToLower(text)
	may := strings.ToUpper(text)

	return min, may
}
