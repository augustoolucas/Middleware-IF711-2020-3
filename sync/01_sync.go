package main

import (
	"fmt"
	"sync"
)

func drive(wg *sync.WaitGroup, m *sync.Mutex, direction string) {
	m.Lock()
	if direction == "right" {
		fmt.Println("carro indo pra direita entrando na pista")
		fmt.Println("carro indo pra direita deixando a pista")
	}
	if direction == "left" {
		fmt.Println("carro indo pra esquerda entrando na pista")
		fmt.Println("carro indo pra esquerda deixando a pista")
	}
	m.Unlock()
	wg.Done()
}

func main() {
	//true -> quando há carros indo para a esquerda
	//false -> quando há carros indo para a direita
	var wg sync.WaitGroup
	var m sync.Mutex
	wg.Add(2)
	go drive(&wg, &m, "left")
	go drive(&wg, &m, "right")
	wg.Wait()
	fmt.Print("Acabou a fila de carros")
}
