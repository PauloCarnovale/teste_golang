package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"
)

const (
	nrPrimos = 10  // número de valores primos para cada magnitude
)

func main() {
	// Informar ao usuário o número de núcleos disponíveis
	numCores := runtime.NumCPU()
	fmt.Printf("Número de núcleos disponíveis: %d\n", numCores)

	// Definir granularidades predefinidas
	granularidades := []int{32, 64, 128, 256, 512, 1024}

	// Definir números de processadores a serem testados
	numProcsList := []int{1, 2, 4, 8, 12, 16, 32}

	var resultados []string

	for _, numValores := range granularidades {
		for _, numProcs := range numProcsList {
			if numProcs > numCores {
				continue // pular se o número de processadores for maior que o disponível
			}
			runtime.GOMAXPROCS(numProcs)

			// Gerar números aleatórios baseados na granularidade
			fmt.Printf("Gerando %d números aleatórios...\n", nrPrimos)
			primos := generateSlice(numValores)

			fmt.Println("****** Valores avaliados: ", primos)
			resSeq := contaPrimosSeq(primos)
			fmt.Printf("\nTempo Seq: %v\n", resSeq)

			end := make(chan int)
			resConc := contaPrimosConc(primos, end)
			fmt.Printf("\nTempo Conc: %v\n", resConc)

			// Garantir que não estamos dividindo por zero
			seqMillis := resSeq.Seconds()
			concMillis := resConc.Seconds()

			var speedup float64
			if concMillis > 0 {
				speedup = seqMillis / concMillis
			} else {
				speedup = 0
			}
			fmt.Printf("Speedup: %f\n", speedup)

			// Armazenar resultados
			resultados = append(resultados, fmt.Sprintf("Granularidade: %d valores, Processadores: %d, Tempo Seq: %v, Tempo Conc: %v, Speedup: %f",
				numValores, numProcs, resSeq, resConc, speedup))
		}
	}

	// Salvando resultados
	saveResults(resultados)
}

func generateSlice(size int) []int {
	slice := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(10000) // valores menores
	}
	return slice
}

func contaPrimosSeq(s []int) time.Duration {
	start := time.Now()
	for i := 0; i < len(s); i++ {
		isPrime(s[i])
	}
	return time.Since(start)
}

func contaPrimosConc(s []int, end chan int) time.Duration {
	start := time.Now()
	for i := 0; i < len(s); i++ {
		go func(i int) {
			isPrime(s[i])
			end <- 1
		}(i)
	}
	for i := 0; i < len(s); i++ {
		<-end
	}
	return time.Since(start)
}

// isPrime verifica se um número é primo
func isPrime(p int) bool {
	if p < 2 {
		return false
	}
	if p%2 == 0 {
		return p == 2
	}
	for i := 3; i*i <= p; i += 2 {
		if p%i == 0 {
			return false
		}
	}
	return true
}

func saveResults(results []string) {
	file, err := os.Create("resultados.txt")
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
		return
	}
	defer file.Close()

	for _, result := range results {
		_, err := file.WriteString(result + "\n")
		if err != nil {
			fmt.Println("Erro ao escrever no arquivo:", err)
			return
		}
	}
	fmt.Println("Resultados salvos em resultados.txt")
}
