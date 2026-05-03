package main

import "fmt"

func main() {

	cores := []string{"Azul", "Verxde", "Amarelo"}

	for i, cores := range cores {
		fmt.Println("A cor no indice", i, ":", cores)
	}

}
