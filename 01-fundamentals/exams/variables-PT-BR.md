# Teste de Avaliação Profissional: Variáveis em Golang


**1. Analise o comportamento do operador de declaração curta (`:=`) em relação ao 'shadowing' (sombreamento) de variáveis. Qual das afirmações abaixo descreve um cenário onde ocorre um erro de compilação em vez de sombreamento ou reatribuição?**
A) Usar `:=` para redeclarar uma variável no mesmo bloco de escopo onde ela já foi declarada, sem introduzir nenhuma variável nova no lado esquerdo.
B) Usar `:=` dentro de um bloco `if` para atribuir um valor a uma variável que possui o mesmo nome de outra declarada no escopo da função externa.
C) Usar `:=` em uma função declarando uma variável local com o mesmo nome de uma variável declarada a nível de pacote (global).
D) Declarar uma nova variável junto com uma existente usando `:=` no mesmo escopo, como por exemplo em `v, err := funcA()`, onde `err` já havia sido declarada.

**2. A alocação de memória para variáveis locais em Go (Stack vs Heap) é decidida pelo compilador através da Análise de Fuga (Escape Analysis). Qual ação envolvendo uma variável local garantirá que o compilador a aloque no Heap em vez de na Stack (pilha)?**
A) Retornar um ponteiro para a variável local a partir da função onde ela foi instanciada.
B) Declarar uma variável de um tipo de estrutura (struct) com dezenas de campos, resultando em um tamanho grande.
C) Passar a variável por valor como argumento para uma função que será executada em uma nova goroutine (usando a palavra-chave `go`).
D) Usar a função `new()` para alocar e inicializar a variável dentro de uma função com ciclo de vida curto e que não retorna a variável.

**3. O valor zero (zero-value) de variáveis que são ponteiros, slices, maps, canais e interfaces é `nil`. No entanto, interfaces podem causar surpresas. Considere que você atribui uma variável de ponteiro de um tipo específico que é nulo (ex: `var p *int = nil`) a uma variável de interface (ex: `var i interface{} = p`). Qual o resultado de validar `if i == nil`?**
A) A comparação `i == nil` resultará em falso (`false`), pois a interface contém um tipo concreto, mesmo que o valor seja nulo.
B) A comparação `i == nil` resultará em verdadeiro (`true`), pois o valor subjacente que a interface armazena aponta para `nil`.
C) Ocorrerá um erro de compilação, impedindo a comparação direta de uma interface vazia com `nil` se um tipo já lhe foi atribuído.
D) A comparação causará um 'panic' em tempo de execução ('nil pointer dereference') ao tentar acessar o valor nulo do ponteiro envelopado.

**4. Como as variáveis são capturadas por closures (funções anônimas) em Go, de maneira padrão, quando criadas dentro de iterações de loops `for` (comportamento histórico antes da correção automática do Go 1.22)?**
A) A closure captura a variável do loop por referência; assim, ao executar as closures após o loop terminar, todas enxergarão o valor final assumido pela variável de iteração.
B) O compilador Go identifica a closure e cria uma cópia local forçada da variável de iteração, garantindo o isolamento dos valores sem intervenção do desenvolvedor.
C) A closure captura a variável pelo seu valor (by value), tornando impossível que uma alteração na variável original dentro do loop afete as funções anônimas geradas.
D) Ocorrerá um 'Data Race' automático reportado em tempo de compilação caso a variável seja capturada e tentada ser usada fora do ciclo sem o uso de um Mutex.

**5. Qual é a diferença fundamental de comportamento em tempo de execução ao declarar uma variável de tipo `map` não inicializada (valor zero) versus declarar uma variável `slice` não inicializada, ao tentar manipular seus valores dinamicamente?**
A) É perfeitamente possível usar a função embutida `append` em um slice `nil` para adicionar novos dados, mas atribuir uma chave e um valor a um map `nil` causará um 'panic'.
B) O valor zero de um `slice` aponta para um array base vazio com capacidade 0 já alocado na memória, ao passo que o `map` nulo não possui alocação, causando panic em ambos os casos se não usarmos a função `make()`.
C) Tentar ler um índice inexistente em um slice `nil` retorna o valor zero do tipo contido, enquanto ler uma chave inexistente em um map `nil` provoca um 'panic'.
D) Variáveis de ambos os tipos, quando atribuídas com seu valor zero, rejeitam leituras e escritas diretas com 'panic' imediatos até que sejam explicitamente alocadas usando o construtor `new`.
