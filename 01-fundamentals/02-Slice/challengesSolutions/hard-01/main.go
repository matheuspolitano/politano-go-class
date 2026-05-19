package main

import "fmt"

func SomaPares(x []int) int {
	var total int
	for _, numero := range x {
		if numero%2 == 0 {
			total = numero + total
		}
	}
	return total
}

func main() {

	list := []int{1, 2, 3, 4, 5}
	fmt.Println(SomaPares(list))
}
