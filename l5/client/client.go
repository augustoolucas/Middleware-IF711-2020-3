package main

import (
	"Middleware-IF711-2020-3/l5/clientproxy"
	"fmt"
	"os"
	"sync"
)

func client(wg *sync.WaitGroup, message string) {
	m := make(map[string]clientproxy.ClientProxy)
	namingService := namingService{m}
	hashing := namingService.Lookup("Hash")
	fmt.Println(hashing.HashPw(message))
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	message := os.Args[1]
	go client(&wg, message)
	wg.Wait()

}

//TODO
