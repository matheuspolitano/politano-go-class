package main

import (
	"fmt"
)

func main() {

	const pi = 3.14
	raio := 5

	newValue := int(pi * float64(raio))

	fmt.Println(newValue)

}
