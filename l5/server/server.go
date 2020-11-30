package main

import (
	"Middleware-IF711-2020-3/l5/clientproxy"
	"Middleware-IF711-2020-3/l5/invoker"
	"Middleware-IF711-2020-3/l5/naming"
	"fmt"
)

// type SRH struct {
// 	ServerHost string
// 	ServerPort int
// }

// func hashRequest(req hashing.Request) string {
// 	hashed := sha256.Sum256([]byte(req.PwRaw))
// 	response := hex.EncodeToString(hashed[:])

// 	//fmt.Println(response)
// 	return response
// }

// // func invoker(receivedReq []byte) []byte {
// // 	request := hashing.Request{PwRaw: ""}
// // 	err := json.Unmarshal(receivedReq, &request)
// // 	shared.ChecaErro(err, "nao foi possivel unmarshal")

// // 	response := hashing.Response{PwSha256: hashRequest(request)}
// // 	responseRaw, err := json.Marshal(response)
// // 	shared.ChecaErro(err, "não foi possível fazer o marshal")

// // 	return responseRaw
// // }

// func server(transportProtocol string) {
// 	for {
// 		if transportProtocol == "TCP" {
// 			l, err := net.Listen("tcp", "localhost:3300")
// 			shared.ChecaErro(err, "nao foi possivel criar servidor tcp")
// 			defer l.Close()

// 			conn, err := l.Accept()
// 			shared.ChecaErro(err, "nao foi possivel aceitar conexao tcp")
// 			defer conn.Close()

// 			receivedReq := make([]byte, 2048)
// 			n, err := conn.Read(receivedReq)

// 			shared.ChecaErro(err, "nao foi possivel receber mensagem do cliente udp")
// 			conn.Write(invoker(receivedReq[:n]))

// 			return
// 		} else if transportProtocol == "UDP" {
// 			addr, err := net.ResolveUDPAddr("udp", ":8030")
// 			conn, err := net.ListenUDP("udp", addr)
// 			shared.ChecaErro(err, "nao foi possivel criar servidor udp")

// 			defer conn.Close()

// 			receivedReq := make([]byte, 2048)
// 			n, addr, err := conn.ReadFromUDP(receivedReq)

// 			shared.ChecaErro(err, "nao foi possivel receber mensagem do cliente udp")

// 			conn.WriteToUDP(invoker(receivedReq[:n]), addr)

// 			return
// 		}
// 	}
// }

func main() {
	m := make(map[string]clientproxy.ClientProxy)
	namingService := naming.NamingService{Table: m}
	hashing := clientproxy.ClientProxy{}

	namingService.Register("Hash", hashing)
	namingService.Register("Add", hashing)

	//invoker
	myInvoker := invoker.ServerInvoker{}
	myInvoker.Invoke()
	fmt.Scanln()
}
