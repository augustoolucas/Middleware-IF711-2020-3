package main

import (
	"fmt"
	"sync"
)

func drive(wg *sync.WaitGroup, direction string) {
	if direction == "right" {
		fmt.Println("carro indo pra direita entrando na pista")
		fmt.Println("carro indo pra direita deixando a pista")
	}
	if direction == "left" {
		fmt.Println("carro indo pra esquerda entrando na pista")
		fmt.Println("carro indo pra esquerda deixando a pista")
	}
	wg.Done()
}

func main() {
	//true -> quando há carros indo para a esquerda
	//false -> quando há carros indo para a direita
	var wg sync.WaitGroup
	wg.Add(2)
	go drive(&wg, "left")
	go drive(&wg, "right")
	wg.Wait()
	fmt.Print("Acabou a fila de carros")
}
