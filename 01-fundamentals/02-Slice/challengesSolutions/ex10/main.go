package main

import "fmt"

func main() {

	aprovados := make([]int, 0, 100)
	notas := []int{7, 4, 9, 3, 8}

	for _, nota := range notas {

		if nota >= 7 {
			aprovados = append(aprovados, nota)
		}

	}
	fmt.Printf("Melhores notas: %d", aprovados)
}
