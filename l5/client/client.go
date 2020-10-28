package main

import (
	"fmt"
	"hashing"
	"os"
	"shared"
	"strings"
	"sync"
)

func client(transportProtocol string, wg *sync.WaitGroup, message string) {
	//fmt.Println(transportProtocol)
	myReq := hashing.Request{PwRaw: message}

	response, err := hashing.HashPw(myReq, strings.ToUpper(transportProtocol))
	shared.ChecaErro(err, "Houve erro na requisição")
	fmt.Println(response)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	transportProtocol := os.Args[1]
	message := os.Args[2]
	go client(transportProtocol, &wg, message)
	wg.Wait()

}

//TODO
