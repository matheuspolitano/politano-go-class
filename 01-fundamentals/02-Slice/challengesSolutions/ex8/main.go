package main

import "fmt"

func main() {

	cores := []string{"Azul", "Verxde", "Amarelo"}

	for i, cor := range cores {
		fmt.Printf("A cor no indice %d é %s\n", i, cor)
	}

}
