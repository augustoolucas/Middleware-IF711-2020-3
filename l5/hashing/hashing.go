package hashing

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"shared"
	"strings"
	"time"
)

//Response em sha256
type Response struct {
	PwSha256 string
}

//Request pro hasher
type Request struct {
	PwRaw string
}

type ClientProxy struct {
	Host     string
	Port     int
	Id       int
	TypeName string
}

type Invocation struct {
	Host    string
	Port    int
	Request Request
}

type Requestor struct{}

func HashPw(message string, transportProtocol string) (string, error) {
	if message == "" {
		err := errors.New("não houve password informado")
		shared.ChecaErro(err, "não houve password informado")
		return "", err
	}

	proxy := ClientProxy{Host: "localhost", Port: 3080, Id: 1, TypeName: "type"}

	// Prepara a invocação ao Requestor
	request := Request{PwRaw: message}
	inv := Invocation{Host: proxy.Host, Port: proxy.Port, Request: request}

	// invoke requestor
	// Invoca o Requestor e aguarda resposta
	req := Requestor{}
	response := req.Invoke(inv, transportProtocol)

	// Envia resposta ao Cliente
	return response.PwSha256, nil
}

func (Requestor) Invoke(inv Invocation, transportProtocol string) Response {
	pwRawBytes, err := json.Marshal(inv.Request)
	shared.ChecaErro(err, "não foi possível fazer o marshal")

	var response = Response{PwSha256: ""}
	err = json.Unmarshal(CRH(pwRawBytes, transportProtocol), &response)
	return response
}

//CRH client request handler
func CRH(pwRawBytes []byte, protocol string) []byte {
	timeoutSeconds := time.Second * 3

	if protocol == "TCP" {
		conn, err := net.DialTimeout(strings.ToLower(protocol), "localhost:3300", timeoutSeconds)
		shared.ChecaErro(err, "nao foi possivel estabelecer conexao tcp")
		defer conn.Close()

		_, err = conn.Write(pwRawBytes)
		shared.ChecaErro(err, "nao foi possivel enviar mensagem tcp")
		response := make([]byte, 2048)
		n, err := conn.Read(response)

		fmt.Println("preso aqui")
		shared.ChecaErro(err, "nao foi possivel receber mensagem tcp")

		return response[:n]
	} else if protocol == "UDP" {
		addr, err := net.ResolveUDPAddr(strings.ToLower(protocol), "localhost:8030")
		shared.ChecaErro(err, "nao foi possivel resolver endereço udp")

		conn, err := net.DialUDP("udp", nil, addr)
		defer conn.Close()

		_, err = conn.Write(pwRawBytes)
		shared.ChecaErro(err, "nao foi possivel enviar mensagem udp")

		response := make([]byte, 2048)
		n, err := conn.Read(response)
		shared.ChecaErro(err, "nao foi possivel receber mensagem udp")

		return response[:n]
	}

	var reponse []byte
	return reponse
}
