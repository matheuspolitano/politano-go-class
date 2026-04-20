## 💻 Desafios de Criação de Código

### Desafio 1: A Troca Mágica (Sem Variável Temporária)
Em muitas linguagens, para trocar o valor de duas variáveis, você precisa criar uma terceira variável "temporária" para guardar o valor enquanto faz a troca. Em Go, isso não é necessário!

**A Tarefa:** Escreva o código dentro da função `main` que:
1. Declare uma variável `x` valendo `10` e uma `y` valendo `20`.
2. Em **uma única linha de código**, inverta os valores (faça `x` valer `20` e `y` valer `10`), sem declarar nenhuma variável nova.
3. Imprima os dois valores.

> **Dica para o aluno:** Lembre-se de como o Go permite lidar com múltiplas variáveis do lado esquerdo e direito do sinal de igual `=` ao mesmo tempo.

---

### Desafio 2: O Choque de Tipos
Go é uma linguagem "fortemente tipada". Isso significa que ela odeia misturar tipos de dados diferentes, a menos que você a force a fazer isso.

**A Tarefa:**
Escreva o código dentro da função `main` que:
1. Declare uma constante `pi` com o valor `3.14`.
2. Declare uma variável `raio` do tipo `int` valendo `5`.
3. Tente multiplicar `pi` por `raio` e guarde o resultado em uma nova variável inteira chamada `areaAproximada`. 
4. Você terá que converter os tipos explicitamente no código para que a compilação funcione sem erros e o resultado perca as casas decimais.

> **Dica para o aluno:** Você pode converter uma variável envolvendo-a no tipo desejado, como `float64(minhaVariavel)` ou `int(minhaVariavel)`. Pense bem onde a conversão deve acontecer para não perder dados antes da multiplicação!

---

## 🛠️ Desafios de Correção de Código (Caça-Bugs)

### Desafio 3: O Fantasma da Redeclaração
O aluno tentou criar um código para gerenciar a pontuação de um jogo, mas o compilador do Go está gritando e se recusando a rodar. 

**A Tarefa:** Identifique os **dois erros** de declaração no código abaixo e conserte-os para que o programa compile corretamente.

```go
package main

import "fmt"

func main() {
    var jogador string
    jogador := "Alex"
    
    pontos := 100
    pontos := 150 
    
    fmt.Println(jogador, "tem", pontos, "pontos")
}
```

> **Dica para o aluno:** Preste muita atenção na diferença entre `var`, `:=` (criar e atribuir) e `=` (apenas atualizar). O compilador não deixa você criar o que já existe no mesmo escopo!

---

### Desafio 4: Matemática Ilegal
O aluno tentou fazer um cálculo simples adicionando um bônus à nota de um teste, mas o código quebrou na linha da soma.

**A Tarefa:** Explique por que o Go não aceitou a linha `notaFinal := nota + bonus` e altere o código para corrigir o problema.

```go
package main

import "fmt"

func main() {
    var nota int = 8
    var bonus float64 = 1.5
    
    notaFinal := nota + bonus 
    
    fmt.Println("A nota final é:", notaFinal)
}
```


