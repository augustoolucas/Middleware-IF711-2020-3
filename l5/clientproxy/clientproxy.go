package clientproxy

import (
	"Middleware-IF711-2020-3/l5/auxiliar"
	"Middleware-IF711-2020-3/l5/requestor"
	"errors"
	"shared"
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
	request := auxiliar.Request{Op: "Hash", Params: params}
	inv := auxiliar.Invocation{Host: proxy.Host, Port: proxy.Port, Request: request}

	// invoke requestor
	// Invoca o Requestor e aguarda resposta
	req := requestor.Requestor{}
	response := req.Invoke(inv).([]interface{})

	// Envia resposta ao Cliente
	return string(response[0].(string)), nil
}

func (ClientProxy) Add(p1 int, p2 int, transportProtocol string) (int, error) {
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
	response := req.Invoke(inv).([]interface{})

	// Envia resposta ao Cliente
	return int(response[0].(int)), nil
}
