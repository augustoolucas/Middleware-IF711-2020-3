package clientproxy

import (
	"Middleware-IF711-2020-3/l5/auxiliar"
	"Middleware-IF711-2020-3/l5/requestor"
	"errors"
	"fmt"
	"shared"
	"strconv"
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
	Request auxiliar.Request
}

func (ClientProxy) HashPw(message string) (string, error) {
	if message == "" {
		err := errors.New("não houve password informado")
		shared.ChecaErro(err, "não houve password informado")
		return "", err
	}

	proxy := ClientProxy{Host: "localhost", Port: 3080, Id: 1, TypeName: "type"}

	// Prepara a invocação ao Requestor
	params := make([]interface{}, 1)
	params[0] = message
	fmt.Println("clientproxy:", params)
	request := auxiliar.Request{Op: "Hash", Params: params}
	inv := auxiliar.Invocation{Host: proxy.Host, Port: proxy.Port, Request: request}

	// invoke requestor
	// Invoca o Requestor e aguarda resposta
	req := requestor.Requestor{}
	response := req.Invoke(inv)

	fmt.Println("clientproxy:", response[0])
	// Envia resposta ao Cliente
	if response[0] == nil {
		return "", nil
	}
	interfaceToString := fmt.Sprintf("%v", response[0])
	fmt.Println("clientproxy:", interfaceToString)
	return interfaceToString, nil
}

func (ClientProxy) Add(p1 int, p2 int) (int, error) {
	proxy := ClientProxy{Host: "localhost", Port: 3080, Id: 1, TypeName: "type"}

	// Prepara a invocação ao Requestor
	params := make([]interface{}, 2)
	params[0] = p1
	params[1] = p2
	request := auxiliar.Request{Op: "Add", Params: params}
	inv := auxiliar.Invocation{Host: proxy.Host, Port: proxy.Port, Request: request}

	// invoke requestor
	// Invoca o Requestor e aguarda resposta
	req := requestor.Requestor{}
	response := req.Invoke(inv)
	if response[0] == nil {
		return 0, nil
	}
	println(response[0])
	interfaceToString := fmt.Sprintf("%v", response[0])
	stringToInt, _ := strconv.Atoi(interfaceToString)
	// Envia resposta ao Cliente
	return stringToInt, nil
}
