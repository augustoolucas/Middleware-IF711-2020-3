package crh

import (
	"fmt"
	"net"
	"shared"
	"strings"
	"time"
)

type CRH struct {
	ServerHost string
	ServerPort int
}

func (CRH) SendReceive(msgToServer []byte, protocol string) []byte {
	timeoutSeconds := time.Second * 3

	if protocol == "TCP" {
		conn, err := net.DialTimeout(strings.ToLower(protocol), "localhost:3300", timeoutSeconds)
		shared.ChecaErro(err, "nao foi possivel estabelecer conexao tcp")
		defer conn.Close()

		_, err = conn.Write(msgToServer)
		fmt.Println("mensagem do cliente", msgToServer)
		shared.ChecaErro(err, "nao foi possivel enviar mensagem tcp")
		response := make([]byte, 2048)
		n, err := conn.Read(response)

		shared.ChecaErro(err, "nao foi possivel receber mensagem tcp")

		return response[:n]
	} else if protocol == "UDP" {
		addr, err := net.ResolveUDPAddr(strings.ToLower(protocol), "localhost:8030")
		shared.ChecaErro(err, "nao foi possivel resolver endereço udp")

		conn, err := net.DialUDP("udp", nil, addr)
		defer conn.Close()

		_, err = conn.Write(msgToServer)
		shared.ChecaErro(err, "nao foi possivel enviar mensagem udp")

		response := make([]byte, 2048)
		n, err := conn.Read(response)
		shared.ChecaErro(err, "nao foi possivel receber mensagem udp")

		return response[:n]
	}

	var reponse []byte
	return reponse
}
