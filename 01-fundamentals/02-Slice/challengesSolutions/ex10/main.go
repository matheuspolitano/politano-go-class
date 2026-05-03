package main

import "fmt"

func main() {

	var aprovados []int
	notas := []int{7, 4, 9, 3, 8}

	for _, nota := range notas {

		if nota >= 7 {
			aprovados = append(aprovados, nota)
		}

	}
	fmt.Println(aprovados)
}
