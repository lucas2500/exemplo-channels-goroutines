package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	// Cria o canal
	winner := make(chan string)

	// Inicia as threads
	for i := 0; i < 10; i++ {
		go ThreadRace(i, winner)
	}

	// Printa a string preenchia pela primeira thread a finalizar a função ThreadRace
	fmt.Println("We have a winner:", <-winner)
}

func ThreadRace(ThreadId int, winner chan string) {

	fmt.Println("Thread", ThreadId, "iniciando...")

	time.Sleep(time.Second * 2)

	winner <- "Thread " + strconv.Itoa(ThreadId)
}
