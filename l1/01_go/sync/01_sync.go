package main

import (
	"fmt"
	"sync"
)

func drive(id int, wg *sync.WaitGroup, m *sync.Mutex, direction string) {
	m.Lock()
	if direction == "right" {
		fmt.Println("carro indo pra direita entrando na pista", id)
		fmt.Println("carro indo pra direita deixando a pista", id)
	}
	if direction == "left" {
		fmt.Println("carro indo pra esquerda entrando na pista", id)
		fmt.Println("carro indo pra esquerda deixando a pista", id)
	}
	m.Unlock()
	wg.Done()
}

func main() {
	//true -> quando há carros indo para a esquerda
	//false -> quando há carros indo para a direita
	var wg sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go drive(i, &wg, &m, "left")
	}
	wg.Wait()
	fmt.Print("Acabou a fila de carros")
}
