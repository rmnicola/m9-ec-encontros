package main

import (
	"fmt"
	"time"
)

// Função que envia uma mensagem para um canal após um sleep de 2 segundos
func printWithDelay(msg string, ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- msg // Envia a mensagem para o canal
}

func main() {
	ch := make(chan string, 3) // Cria um canal com buffer para 3 strings

	start := time.Now() // Início da medição do tempo

	// Chamando a função 3 vezes de forma concorrente
	go printWithDelay("Mensagem 1", ch)
	go printWithDelay("Mensagem 2", ch)
	go printWithDelay("Mensagem 3", ch)

	// Recebe as mensagens do canal e imprime
	for i := 0; i < 3; i++ {
		msg := <-ch
		fmt.Println(msg)
	}

	close(ch)

	elapsed := time.Since(start) // Tempo total de execução
	fmt.Printf("Tempo total de execução: %s\n", elapsed)
}
