package main

import (
	"fmt"
	"sync"
	"time"
)

// Função que exibe uma mensagem após um sleep de 2 segundos
func printWithDelay(msg string, wg *sync.WaitGroup) {
	defer wg.Done() // Indica que a goroutine foi concluída
	time.Sleep(2 * time.Second)
	fmt.Println(msg)
}

func main() {
	var wg sync.WaitGroup // Cria um WaitGroup

	start := time.Now() // Início da medição do tempo

	// Adiciona 3 ao contador de tarefas do WaitGroup
	wg.Add(3)

	// Chamando a função 3 vezes de forma concorrente
	go printWithDelay("Mensagem 1", &wg)
	go printWithDelay("Mensagem 2", &wg)
	go printWithDelay("Mensagem 3", &wg)

	wg.Wait() // Aguarda todas as goroutines serem concluídas

	elapsed := time.Since(start) // Tempo total de execução
	fmt.Printf("Tempo total de execução: %s\n", elapsed)
}
