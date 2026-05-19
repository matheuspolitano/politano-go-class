package main

import "fmt"

// Exercício 3: Nível Difícil Desenvolva uma função chamada AgruparEmLotes que
// receba dois parâmetros: um slice de inteiros e um tamanho de lote (um número inteiro tamanho).
//  A função deve dividir o slice original em vários slices menores,
//  cada um com a capacidade máxima do tamanho passado,
//  e retornar um slice de slices de inteiros ([][]int).
// Atenção aos detalhes: você precisará lidar com o cálculo
// correto dos índices para não gerar um erro de out of bounds
// (acessar posições que não existem), principalmente no último lote,
// que frequentemente será menor que o tamanho estipulado.

// input: todos os item, tamanho do lote
// output:   slice de slices de inteiros ([][]int)

////.  [[1,2,3],[4,5,6],[1,1,1]]
////   [1,2,3]
//// Jagged Slices vamos usar
//// Flat Slices n pode ser usado

func AgruparEmLotes(sliceGrande []int, tamanhoLote int) [][]int {
	// tamanhoLote = 4
	// len(sliceGrande) = 10
	// (4+10-1)/4
	quantidadeLote := (tamanhoLote + len(sliceGrande) - 1) / tamanhoLote
	listaLotes := make([][]int, quantidadeLote)
	for i := 0; i < quantidadeLote; i++ {
		fmt.Println(i)
		listaLotes[i] = make([]int, 0, tamanhoLote)
	}

	for i, itemLista := range sliceGrande {
		// i = 0
		// itemLista = cego

		// // experado: listaLotes[2] = append(listaLotes[2], itemLista)
		// listaLotes[10/4] = append(listaLotes[10/4], itemLista)
		// // resulado sem o adiconar: listaLotes[3] = append(listaLotes[3], itemLista)

		// experado i = 0
		// tamanhoLote = 4
		// indexLote := (0 / 4)
		// experado i = 10
		// tamanhoLote = 4
		// indexLote := (10  / 4)
		indexLote := i / tamanhoLote
		// fmt.Println(i)
		listaLotes[indexLote] = append(listaLotes[indexLote], itemLista)

	}
	return listaLotes
}

func main() {
	fmt.Println(AgruparEmLotes([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 4))
}
