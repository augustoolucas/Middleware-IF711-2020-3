package clientproxy

import (
	"Middleware-IF711-2020-3/l5/hashing"
	"fmt"
)

type ClientProxy struct {
	Host     string
	Port     int
	Id       int
	TypeName string
}

type Invocation struct {
	Host    string
	Port    int
	Request hashing.Request
}

func (proxy ClientProxy) Hash(rawString string) string {
	params := make([]interface{}, 1)
	params[0] = rawString
	request := hashing.Request{PwRaw: rawString}

	invok := Invocation{proxy.Host, proxy.Port, request}

	fmt.Println("Antes da invocacao: ", invok)
	//da pro invoker o invok e retorna com a resposta do cliente
	return ""
}
