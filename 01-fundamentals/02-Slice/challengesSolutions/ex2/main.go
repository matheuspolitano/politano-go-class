package main

import "fmt"

func main() {

	lista := []string{"Maca", "Banana", "Laranja"}

	fmt.Println(lista[1])

	lista[0] = "Morango"
	fmt.Println(lista[0])
}
