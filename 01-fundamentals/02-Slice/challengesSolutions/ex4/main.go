package main

import "fmt"

func main() {

	numeros := []int{1, 2}
	numeros = append(numeros, 3, 4, 5)

	fmt.Println(len(numeros))
}
