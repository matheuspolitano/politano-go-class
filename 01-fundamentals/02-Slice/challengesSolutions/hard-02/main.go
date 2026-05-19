package main

import "fmt"

func removerDuplicados(pessoas []string) []string {
	dentroFesta := make([]string, 0, len(pessoas))

	// NAO PRECISA DO IF == 0 PORUQE: pula esse loop inteiro se a lista for vazia

	for _, pessoaFora := range pessoas {
		podeEntrar := true

		for _, pessoaDentro := range dentroFesta {
			if pessoaDentro == pessoaFora {
				podeEntrar = false
				break
			}
		}

		if podeEntrar {
			dentroFesta = append(dentroFesta, pessoaFora)
		}
	}

	return dentroFesta
}

func main() {
	fmt.Println(removerDuplicados([]string{"a", "b", "a", "b", "a", "gabriel", "SAdas"}))
}
