# Gabarito e Explicações: Variáveis em Golang

---

### Questão 1
* **A) [CORRETA]** O operador `:=` exige que pelo menos uma variável no lado esquerdo seja nova no escopo atual. Se todas já existirem no mesmo bloco, o compilador lança o erro *'no new variables on left side of :='*.
* **B) [INCORRETA]** Neste caso, não há erro de compilação. Como o bloco `if` cria um novo escopo, ocorre o sombreamento (shadowing) da variável do escopo externo.
* **C) [INCORRETA]** Isso é totalmente permitido. A variável local simplesmente sombreia a variável de nível de pacote, não gerando nenhum erro.
* **D) [INCORRETA]** Essa é uma característica permitida e muito comum em Go (ex: tratamento de erros). A variável já existente (`err`) sofre uma reatribuição, enquanto a nova (`v`) é declarada normalmente.

### Questão 2
* **A) [CORRETA]** Se o endereço de uma variável local for retornado, sua referência 'escapa' do escopo da função. Para que a memória não seja corrompida quando a *Stack Frame* da função for destruída, o compilador transfere a alocação dessa variável para o Heap.
* **B) [INCORRETA]** O tamanho por si só não garante a alocação no Heap, a menos que ultrapasse os limites máximos da Stack (o que é raro para structs comuns que não escapam do escopo).
* **C) [INCORRETA]** Passar por valor significa que uma cópia dos dados é feita e entregue à nova goroutine. A variável original não escapa de sua Stack original por causa dessa cópia.
* **D) [INCORRETA]** Em Go, a palavra-chave `new()` não obriga a alocação no Heap (diferente de linguagens como C++). Se o compilador provar que a referência não escapa do escopo local, ele vai otimizar e alocar a variável na Stack.

### Questão 3
* **A) [CORRETA]** Uma variável de interface por baixo dos panos é uma tupla `(tipo, valor)`. Uma interface só é considerada estritamente `nil` se ambos (tipo e valor) forem nulos. Ao atribuir um `*int` nulo, a interface passa a armazenar o tipo `*int`, o que faz com que a comparação com um `nil` puro retorne `false`.
* **B) [INCORRETA]** Esta é uma das armadilhas mais famosas do Go. A interface não olha apenas para o valor que ela envelopa; ela verifica sua própria estrutura de metadados. Como ela possui um tipo associado, ela não é `nil`.
* **C) [INCORRETA]** A comparação entre qualquer interface e `nil` é válida sintaticamente e perfeitamente legal, não gerando erros de compilação.
* **D) [INCORRETA]** A operação de comparação `== nil` avalia apenas os metadados da interface (a tupla). Ela não tenta desreferenciar o ponteiro que está lá dentro, portanto não causa 'panic'.

### Questão 4
* **A) [CORRETA]** Historicamente (antes do Go 1.22), a variável de loop `for` era alocada apenas uma vez na memória. As closures iterativas fechavam-se sobre esse mesmo endereço de memória. Quando o loop terminava e as closures eram executadas, todas liam o valor final que ficou na variável.
* **B) [INCORRETA]** Este passou a ser o comportamento padrão *apenas a partir do Go 1.22*. Historicamente, o desenvolvedor precisava forçar essa cópia local manualmente (ex: `v := v` dentro do loop).
* **C) [INCORRETA]** Closures em Go sempre capturam variáveis do ambiente externo por referência (ponteiro implícito), nunca copiando o valor no momento da declaração.
* **D) [INCORRETA]** O Go não detecta "Data Races" em tempo de compilação. O compilador não te impede de cometer esse erro lógico. Para detectar corridas de dados, é necessário rodar o código ou testes com a flag `-race`.

### Questão 5
* **A) [CORRETA]** A função `append` é inteligente: se o slice for `nil`, ela aloca um novo array base e retorna o novo slice. No entanto, um `map` nulo não possui sua estrutura de hash interna inicializada; tentar atribuir um valor a ele resulta imediatamente em *'panic: assignment to entry in nil map'*.
* **B) [INCORRETA]** O valor zero de um slice é estritamente `nil` (sem nenhum array alocado, tamanho 0 e capacidade 0). Ele não aponta para um array vazio pré-alocado.
* **C) [INCORRETA]** É o oposto: consultar (apenas leitura) uma chave em um map `nil` funciona perfeitamente e retorna o valor zero do tipo do valor. Porém, tentar acessar um índice num slice `nil` (ex: `slice[0]`) causa um *'panic: index out of range'*.
* **D) [INCORRETA]** Como explicado na opção A e C, existem operações permitidas em slices nulos (como `append`) e em maps nulos (como leituras). Nem tudo resulta em panic. Além disso, a inicialização padrão em Go para esses tipos é feita com `make`, e não com `new`.
