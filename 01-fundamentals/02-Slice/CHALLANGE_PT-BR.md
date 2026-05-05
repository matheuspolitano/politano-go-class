## 💻 Desafios de Criação de Código (Nível Iniciante)

### Desafio 1: O Slice Fantasma
Em Go, declarar um slice sem valores o deixa em um estado chamado `nil`. É como ter uma caixa que ainda não existe fisicamente.

**A Tarefa:** Escreva o código dentro da função `main` que:
1. Declare um slice de inteiros chamado `vazio` sem atribuir nenhum valor inicial (use apenas `var`).
2. Imprima o slice e, logo após, imprima o seu tamanho usando a função `len()`.

> **Dica:** Slices vazios em Go são seguros de usar com `len()`, resultando em 0.

---

### Desafio 2: A Fruta Secreta
Slices funcionam como listas onde cada item tem um índice, começando sempre do zero.

**A Tarefa:**
1. Crie um slice literal chamado `frutas` contendo: `"Maçã"`, `"Banana"` e `"Laranja"`.
2. Imprima **apenas** o segundo elemento da lista.
3. Altere o valor de `"Maçã"` para `"Morango"` e imprima o slice completo.

---

### Desafio 3: O Passageiro Extra
Diferente de arrays, os slices podem "crescer". A função `append` é a ferramenta mágica para isso.

**A Tarefa:**
1. Crie um slice de strings chamado `onibus` com os nomes `"Ana"` e `"Beto"`.
2. Use a função `append` para adicionar a `"Clara"` ao ônibus.
3. Imprima o resultado final.

---

### Desafio 4: A Caravana de Números
O `append` aceita múltiplos argumentos de uma só vez.

**A Tarefa:**
1. Crie um slice chamado `numeros` que comece com `1, 2`.
2. Em uma **única linha de código** usando `append`, adicione os números `3, 4, 5`.
3. Imprima o tamanho (`len`) final desse slice.

---

### Desafio 5: Corte de Precisão (Slicing)
Podemos criar sub-listas definindo onde começar e onde parar.

**A Tarefa:**
Dado o slice: `letras := []string{"A", "B", "C", "D", "E"}`
1. Crie um novo slice chamado `miolo` que contenha apenas `"B", "C", "D"`.
2. Imprima o slice `miolo`.

---

### Desafio 6: O Reservatório de Dados (Make)
Usamos `make` para preparar um slice com capacidade extra.

**A Tarefa:**
1. Use `make` para criar um slice de inteiros com **tamanho 2** e **capacidade 5**. Chame-o de `deposito`.
2. Imprima o slice, o seu `len()` e o seu `cap()`.

---

### Desafio 7: A Fila Anda (Shift)
Para remover o primeiro elemento, "fatiamos" o slice do índice 1 em diante.

**A Tarefa:**
Dado o slice: `fila := []string{"Pessoa1", "Pessoa2", "Pessoa3"}`
1. "Remova" a `"Pessoa1"` reatribuindo à variável `fila` o corte `fila[1:]`.
2. Imprima a fila resultante.

---

### Desafio 8: O Passeio Completo (For Range)
O Go possui uma estrutura especial e muito elegante chamada `range`. Ela facilita muito percorrer todos os itens de um slice, entregando de mão beijada a posição (índice) e o item daquela posição (valor).

**A Tarefa:**
1. Crie um slice de strings chamado `cores` contendo `"Azul"`, `"Verde"` e `"Amarelo"`.
2. Use um loop `for indice, cor := range cores` para percorrer a lista.
3. Dentro do loop, imprima uma mensagem para cada item informando sua posição e nome (exemplo: `"A cor no índice 0 é Azul"`).

---

### Desafio 9: A Soma do Tesouro (Acumulador)
Muitas vezes precisamos passar por todos os itens de um slice para calcular um resultado final, como o total de uma compra ou a soma de pontos. O `for` é essencial para isso.

**A Tarefa:**
1. Crie um slice de inteiros chamado `moedas` com os valores `5, 10, 25, 50`.
2. Declare uma variável inteira chamada `total` com o valor inicial de `0`.
3. Use um loop `for` com `range` para somar cada moeda à variável `total`.
4. Após o loop terminar, imprima o valor final de `total`.

> **Dica:** Se você usar o `range` e não precisar do índice, lembre-se de usar o *blank identifier* (`_`) para ignorá-lo: `for _, moeda := range moedas`.

---

### Desafio 10: A Seleção Rigorosa (Filtro)
Podemos combinar a repetição de um `for`, a tomada de decisão de um `if` e o crescimento de um slice com `append` para filtrar dados.

**A Tarefa:**
Dado o slice: `notas := []int{7, 4, 9, 3, 8}`
1. Crie um novo slice de inteiros (vazio) chamado `aprovados`.
2. Use um loop `for` para percorrer todas as `notas`.
3. Dentro do loop, faça uma verificação (`if`): se a nota for **maior ou igual a 7**, adicione-a ao slice `aprovados` usando `append`.
4. Fora do loop, imprima o slice `aprovados` para ver apenas as notas altas.


### Desafio 11: A Grande Ida ao Cinema (Capacidade e Realocação)

Você e seus  amigos  decidiram ir ao cinema para assistir a uma estreia muito aguardada. A regra de ouro do grupo é clara: **todos precisam sentar um do lado do outro na mesma fileira** (exatamente como os elementos de um *slice* em Go, que precisam de memória contígua).

Como vocês não sabem exatamente quantos amigos vão aparecer, começam sentando na primeira fileira vazia que encontram. Conforme novos amigos chegam, se a fileira atual não tiver mais cadeiras vazias ao lado de vocês (ou seja, se o limite da reserva for atingido), **todo o grupo** precisa levantar, recolher suas pipocas e se mudar para uma fileira maior!

**O Desafio:**
Desenvolva um programa em Go para simular a chegada de amigos ao cinema e descubra o quão exaustivo foi para o grupo ficar mudando de lugar.

**Requisitos do Código:**
1. Declare um *slice* de *strings* chamado `amigos` (comece sem usar a função `make`, apenas declare a variável nula/vazia).
2. Crie uma variável para rastrear a `capacidadeAnterior` do *slice* e outra para contar o número de `mudancasDeFileira`.
3. Crie um loop `for` que simule a chegada de **25 amigos** ao cinema. A cada repetição, adicione um novo amigo ao *slice* usando a função `append()` (ex: "Amigo 1", "Amigo 2"...).
4. A cada novo amigo adicionado, verifique se a **capacidade** do *slice* (`cap`) mudou em relação à `capacidadeAnterior`.
5. Se a capacidade mudou, significa que o grupo teve que se mudar! Imprima uma mensagem na tela no formato: 
   *"A fileira encheu! O grupo que agora tem X pessoas precisou se mudar para uma fileira com Y cadeiras."*
6. No final do programa, exiba um resumo contendo:
   * O total de amigos que chegaram (usando `len`).
   * A quantidade final de cadeiras da última fileira (usando `cap`).
   * O total de vezes que o grupo precisou trocar de lugar.


**Exemplo de Saída Esperado antes de fazer o Bônus:**
```text
A fileira encheu! O grupo que agora tem 1 pessoas precisou se mudar para uma fileira com 1 cadeiras.
A fileira encheu! O grupo que agora tem 2 pessoas precisou se mudar para uma fileira com 2 cadeiras.
A fileira encheu! O grupo que agora tem 3 pessoas precisou se mudar para uma fileira com 4 cadeiras.
A fileira encheu! O grupo que agora tem 5 pessoas precisou se mudar para uma fileira com 8 cadeiras.
...

--- RESUMO DA SESSÃO ---
Total de amigos sentados (Length): 25
Tamanho da fileira final (Capacity): 32
Total de mudanças de fileira: 6 vezes.
```
```