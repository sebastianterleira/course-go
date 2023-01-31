package main

import "fmt"

func main() {
	/* nums := []int{1, 4, 59, 5, 81, 34}
	result := filter(nums, func(num int) bool {
		if num <= 10 {
			return true
		} else {
			return false
		}
	})
	fmt.Println(result) */

	x := hello("Sebastian")
	fmt.Println(x("Terleira"))
}

func hello(name string) func(string) string {
	return func(text string) string {
		return "Hello " + name + " " + text 
	}
}


/* func filter(nums []int, callback func(int) bool) []int {
	result := []int{}
	for _, v := range nums {
		if callback(v) {
			result = append(result, v)
		} 
	}

	return result
} */