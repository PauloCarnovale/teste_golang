# Teste de Performance com Cálculo de Números Primos

Este projeto tem como objetivo medir a performance do cálculo de números primos utilizando diferentes granularidades e números de processadores. Os resultados obtidos são usados para analisar o speedup obtido com a execução concorrente em comparação com a execução sequencial.

## Descrição do Código

O código está escrito em Go e consiste em gerar números aleatórios, verificar se são primos e medir o tempo de execução das verificações de forma sequencial e concorrente. O código avalia o desempenho usando diferentes granularidades (número de valores) e diferentes números de processadores.

### Estrutura do Código

1. **Constantes**
    - `nrPrimos`: Número de valores primos para cada magnitude.

2. **Função `main`**
    - Obtém o número de núcleos disponíveis na máquina.
    - Define granularidades predefinidas: 32, 64, 128, 256, 512 e 1024.
    - Define os números de processadores a serem testados: 1, 2, 4, 8, 12, 16, 32.
    - Para cada combinação de granularidade e número de processadores:
        - Define o número de processadores a serem usados.
        - Gera números aleatórios com base na granularidade.
        - Mede o tempo de execução sequencial.
        - Mede o tempo de execução concorrente.
        - Calcula o speedup.
        - Armazena os resultados.
    - Salva os resultados em um arquivo `resultados.txt`.

3. **Função `generateSlice`**
    - Gera um slice de números aleatórios com tamanho definido pela granularidade.

4. **Função `contaPrimosSeq`**
    - Verifica se cada número no slice é primo de forma sequencial.
    - Retorna o tempo de execução.

5. **Função `contaPrimosConc`**
    - Verifica se cada número no slice é primo de forma concorrente utilizando goroutines.
    - Retorna o tempo de execução.

6. **Função `isPrime`**
    - Verifica se um número é primo.

7. **Função `saveResults`**
    - Salva os resultados obtidos em um arquivo `resultados.txt`.

### Como Executar

1. **Compile o programa**:
    ```sh
    go build -o teste_golang
    ```

2. **Execute o programa compilado**:
    ```sh
    ./teste_golang
    ```

### Resultados

Os resultados serão salvos em um arquivo chamado `resultados.txt`, contendo informações sobre a granularidade, número de processadores, tempos de execução sequenciais e concorrentes, e o speedup calculado. Esses dados podem ser usados para criar gráficos e analisar a performance do código.

### Dependências

- [Go](https://golang.org/) instalado em sua máquina.

### Contribuições

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests.

### Licença

Este projeto está licenciado sob a Licença MIT. Veja o arquivo LICENSE para mais detalhes.
