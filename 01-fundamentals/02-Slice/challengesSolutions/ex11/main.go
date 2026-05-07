// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	var amigos []string

	capacidadeAnterior := 0
	mudancasDeFileira := 0

	for i := 1; i <= 33; i++ {

		amigos = append(amigos, "amigo-xyz")

		capacidadeAtual := cap(amigos)

		if capacidadeAtual > capacidadeAnterior {

			fmt.Printf("A fileira encheu! O grupo agora tem %d pessoas, e se mudaram para uma fileira com %d lugares \n", len(amigos), capacidadeAtual)

			capacidadeAnterior = capacidadeAtual

			mudancasDeFileira++

		}

	}

	//RESUMAO

	fmt.Printf("Total de amigos sentados: %d\n", len(amigos))
	fmt.Printf("Tamanho da fileira final: %d\n", cap(amigos))
	fmt.Printf("Total de mudanca de fileira: %d vezez \n", mudancasDeFileira)

}
