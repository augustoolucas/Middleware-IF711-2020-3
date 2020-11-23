package main

import (
	"Middleware-IF711-2020-3/l5/clientproxy"
	"Middleware-IF711-2020-3/l5/naming"
	"fmt"
	"os"
	"sync"
)

func client(wg *sync.WaitGroup, message string) {
	m := make(map[string]clientproxy.ClientProxy)
	namingService := naming.NamingService{Table: m}
	hashing := namingService.Lookup("Hash")
	adding := namingService.Lookup("Add")

	for i := 0; i < 1; i++ {
		response, _ := hashing.HashPw(message)
		fmt.Println("Server response:", response)
		response2, _ := adding.Add(1, 2)
		fmt.Println("Server response:", response2)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	message := os.Args[1]
	fmt.Println("Message:", message)

	wg.Add(1)
	go client(&wg, message)

	wg.Wait()
}
