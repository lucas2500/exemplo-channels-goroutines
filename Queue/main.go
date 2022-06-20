package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	// Cria o canal
	message := make(chan string)

	for i := 0; i < 10; i++ {

		// Inicia thread de processamento da mensagem
		go ReadMessage(message)

		// Adiciona conteúdo da mensagem no canal
		message <- "Processando mensagem: " + strconv.Itoa(i)

		// Aguarda 1 segundo após adicionar a mensagem no canal
		time.Sleep(time.Second * 1)
		// Dado que o scheduler das goroutines não é preemptivo,
		// é necessário fazer a thread principal "dormir" para que
		// as demais tenham tempo de serem executadas.
		// Sem o sleep a thread principal monopoliza o uso da CPU
		// e finaliza o programa após sua execução, não dando chance
		// de nenuma outra thread ser executada
	}
}

func ReadMessage(message chan string) {
	fmt.Println(<-message)
}
