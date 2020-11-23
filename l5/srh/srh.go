package srh

import (
	"net"
	"shared"
)

type SRH struct {
	ServerHost string
	ServerPort int
}

var l net.Listener
var conn net.Conn
var err error

// func hashingRemote(receivedReq []byte) []byte {
// 	request := hashing.Request{Op: "", Params: []interface{}{}}
// 	err := marshaller.Unmarshal(receivedReq, &request)
// 	shared.ChecaErro(err, "nao foi possivel unmarshal")

// 	if request.Op == "Hash" {
// 		response := hashing.Response{PwSha256: hashRequest(request)}
// 		responseRaw, err := marshaller.Marshal(response)
// 	}
// 	shared.ChecaErro(err, "não foi possível fazer o marshal")

// 	return responseRaw
// }

func (srh SRH) Receive() []byte {
	transportProtocol := shared.TRANSPORT_PROTOCOL
	reponse := make([]byte, 1)
	if transportProtocol == "TCP" {
		l, err = net.Listen("tcp", "localhost:3300")
		shared.ChecaErro(err, "nao foi possivel criar servidor tcp")

		conn, err = l.Accept()
		shared.ChecaErro(err, "nao foi possivel aceitar conexao tcp")

		receivedReq := make([]byte, 2048)
		n, err := conn.Read(receivedReq)

		return receivedReq[:n]

		shared.ChecaErro(err, "nao foi possivel receber mensagem do cliente tcp")
	} else {
		addr, err := net.ResolveUDPAddr("udp", ":8030")
		conn, err := net.ListenUDP("udp", addr)
		shared.ChecaErro(err, "nao foi possivel criar servidor udp")

		defer conn.Close()

		receivedReq := make([]byte, 2048)
		n, addr, err := conn.ReadFromUDP(receivedReq)

		shared.ChecaErro(err, "nao foi possivel receber mensagem do cliente udp")
		return receivedReq[:n]

	}

	return reponse
}

func (srh SRH) Send(msgBytes []byte) {
	transportProtocol := shared.TRANSPORT_PROTOCOL
	
	if transportProtocol == "TCP" {
		_, err := conn.Write(msgBytes)
		shared.ChecaErro(err, "server: nao foi possivel enviar mensagem tcp")
	}

	conn.Close()
	l.Close()
}
