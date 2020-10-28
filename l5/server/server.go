package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hashing"
	"net"
	"os"
	"shared"
	"strings"
)

func hashRequest(req hashing.Request) string {
	var hashed [32]byte
	hashed = sha256.Sum256([]byte(req.PwRaw))
	response := hex.EncodeToString(hashed[:])

	fmt.Println(response)
	return response
}

func invoker(receivedReq []byte) []byte {
	request := hashing.Request{PwRaw: ""}
	err := json.Unmarshal(receivedReq, &request)
	shared.ChecaErro(err, "nao foi possivel unmarshal")

	response := hashing.Response{PwSha256: hashRequest(request)}
	responseRaw, err := json.Marshal(response)
	shared.ChecaErro(err, "não foi possível fazer o marshall")

	return responseRaw
}

func server(transportProtocol string) {
	for {
		if transportProtocol == "TCP" {
			fmt.Println("aqui")

			l, err := net.Listen("tcp", "localhost:3300")
			shared.ChecaErro(err, "nao foi possivel criar servidor tcp")
			fmt.Println("aqui2")
			defer l.Close()

			conn, err := l.Accept()
			fmt.Println("aqui3")
			shared.ChecaErro(err, "nao foi possivel aceitar conexao tcp")
			defer conn.Close()

			//parse na resposta TCP
			//receivedReq, _, err := bufio.NewReader(conn).ReadLine()
			receivedReq := make([]byte, 2048)
			n, err := conn.Read(receivedReq)

			fmt.Println("aqui4")
			shared.ChecaErro(err, "nao foi possivel receber mensagem do cliente")
			conn.Write(invoker(receivedReq[:n]))
			fmt.Println("Sucesso!")

			return
		} else if transportProtocol == "UDP" {
			fmt.Println("aqui")

			addr, err := net.ResolveUDPAddr("udp", ":8030")
			conn, err := net.ListenUDP("udp", addr)
			fmt.Println("aqui2")

			shared.ChecaErro(err, "nao foi possivel criar servidor udp")

			defer conn.Close()

			receivedReq := make([]byte, 2048)
			n, addr, err := conn.ReadFromUDP(receivedReq)
			fmt.Println("aqui3")

			shared.ChecaErro(err, "nao foi possivel receber mensagem do cliente")

			conn.WriteToUDP(invoker(receivedReq[:n]), addr)
			fmt.Println("aqui4")

			fmt.Println("Sucesso!")

			return
		}
	}
}

func main() {
	transportProtocol := os.Args[1]
	go server(strings.ToUpper(transportProtocol))
	fmt.Scanln()

}

//TODO
