package main

import "fmt"

func main() {

	numeros := make([]int, 2, 5)

	fmt.Println(len(numeros))
	fmt.Println(cap(numeros))
}
