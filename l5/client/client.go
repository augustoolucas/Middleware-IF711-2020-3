package main

import (
	"Middleware-IF711-2020-3/l5/auxiliar"
	"Middleware-IF711-2020-3/l5/clientproxy"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"shared"
	"sync"
	"time"
)

func lookup(wg2 *sync.WaitGroup, str string) clientproxy.ClientProxy {
	conn, err := net.DialTimeout("tcp", "localhost:3315", 3*time.Second)

	//myProxy não usado no lado do cliente
	//param opicional
	myProxy := auxiliar.Proxy{"", 0, 0, ""}
	msgBytes, _ := json.Marshal(auxiliar.RequestNaming{"Lookup", str, myProxy})

	_, err = conn.Write(msgBytes)
	shared.ChecaErro(err, "nao foi possivel enviar mensagem tcp")

	response := make([]byte, 2048)
	n, err := conn.Read(response)

	shared.ChecaErro(err, "nao foi possivel receber mensagem tcp")

	json.Unmarshal(response[:n], &myProxy)

	clientProxy := clientproxy.ClientProxy{
		Host:     myProxy.Host,
		Port:     myProxy.Port,
		Id:       myProxy.Id,
		TypeName: myProxy.TypeName,
	}
	fmt.Println(clientProxy)
	conn.Close()
	wg2.Done()
	return clientProxy
}

var hasher = make(map[string]clientproxy.ClientProxy)

func client(wg *sync.WaitGroup, mtx *sync.Mutex, message string) {

	//namingService := naming.NamingService{Table: m}

	//mtx.Lock()
	var wg2 sync.WaitGroup
	wg2.Add(1)
	if _, ok := hasher["Hasher"]; !ok {
		fmt.Println("aqui")
		hasher["Hasher"] = lookup(&wg2, "Hasher")
	}
	wg2.Wait()
	//mtx.Unlock()
	start := time.Now()
	for i := 0; i < 10; i++ {
		response, _ := hasher["Hasher"].HashSha256(message)
		fmt.Println("Server response:", response)
		fmt.Println("Client: #", i+1)
		//response2, _ := adding.Add(1, 2)
		//fmt.Println("Server response:", response2)
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	message := os.Args[1]
	var mtx sync.Mutex
	fmt.Println("Message:", message)
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go client(&wg, &mtx, message)
	}

	wg.Wait()
}
