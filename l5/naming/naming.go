package main

import (
	"Middleware-IF711-2020-3/l5/auxiliar"
	"encoding/json"
	"fmt"
	"net"
	"shared"
)

type NamingService struct {
	Table map[string]auxiliar.Proxy
}

// Register registra o clientProxy no serviço de nomes
func (naming *NamingService) Register(name string, proxy auxiliar.Proxy) bool {
	naming.Table[name] = auxiliar.Proxy{
		proxy.Host,
		proxy.Port,
		proxy.Id,
		proxy.TypeName,
	}
	_, found := naming.Table[name]
	return found
}

// Lookup retorna o clientproxy do registro de nomes
func (naming NamingService) Lookup(name string) auxiliar.Proxy {
	value, _ := naming.Table[name]
	return value
}

// List retorna o map do serviço de nomes
func (naming NamingService) List() map[string]auxiliar.Proxy {
	return naming.Table
}

var l net.Listener
var conn net.Conn
var err error

func main() {
	m := make(map[string]auxiliar.Proxy)
	namingService := NamingService{Table: m}

	for {
		objRecv := auxiliar.RequestNaming{}
		fmt.Println("Listening...")
		for {
			l, err = net.Listen("tcp", "localhost:3315")
			if err == nil {
				break
			}
		}
		//shared.ChecaErro(err, "nao foi possivel criar servidor tcp")
		//shared.ChecaErro(err, "naming.go")

		for {
			conn, err = l.Accept()
			if err == nil {
				break
			}
		}

		receivedReq := make([]byte, 2048)
		n, err := conn.Read(receivedReq)

		shared.ChecaErro(err, "naming.go")

		//fmt.Println(objRecv)

		err = json.Unmarshal(receivedReq[:n], &objRecv)
		if objRecv.Op == "Register" {
			fmt.Println("Register...")
			namingService.Register(objRecv.Arg, objRecv.Proxy)
		}
		if objRecv.Op == "Lookup" {
			fmt.Println("Lookup...")
			ret := namingService.Lookup(objRecv.Arg)
			msgBytes, _ := json.Marshal(ret)
			_, err := conn.Write(msgBytes)
			shared.ChecaErro(err, "naming.go")
			fmt.Println("message sent to client")
		}
		conn.Close()
		l.Close()
	}
}
