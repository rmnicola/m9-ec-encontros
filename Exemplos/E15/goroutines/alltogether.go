package main

import (
	"fmt"
	"sync"
	"time"
)

// Função que envia uma mensagem para um canal após um sleep de 2 segundos
func printWithDelay(msg string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done() // Indica que a goroutine foi concluída
	time.Sleep(2 * time.Second)
	ch <- msg // Envia a mensagem para o canal
}

func main() {
	var wg sync.WaitGroup   // Cria um WaitGroup
	ch := make(chan string) // Cria um canal sem buffer

	start := time.Now() // Início da medição do tempo

	// Adiciona 3 ao contador de tarefas do WaitGroup
	wg.Add(3)

	// Chamando a função 3 vezes de forma concorrente
	go printWithDelay("Mensagem 1", ch, &wg)
	go printWithDelay("Mensagem 2", ch, &wg)
	go printWithDelay("Mensagem 3", ch, &wg)

	// Goroutine que aguarda todas as outras goroutines serem concluídas e processa as mensagens do canal
	go func() {
		defer close(ch) // Fecha o canal para indicar que não há mais mensagens a serem enviadas
		for msg := range ch {
			fmt.Println(msg)
		}
	}()

	wg.Wait() // Aguarda a goroutine de processamento concluir

	elapsed := time.Since(start) // Tempo total de execução
	fmt.Printf("Tempo total de execução: %s\n", elapsed)
}
