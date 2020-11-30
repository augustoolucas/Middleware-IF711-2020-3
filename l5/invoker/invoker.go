package invoker

import (
	"Middleware-IF711-2020-3/l5/auxiliar"
	"Middleware-IF711-2020-3/l5/calculator"
	"Middleware-IF711-2020-3/l5/marshaller"
	"Middleware-IF711-2020-3/l5/miop"
	"Middleware-IF711-2020-3/l5/pooling"
	"Middleware-IF711-2020-3/l5/srh"
	"encoding/json"
	"fmt"
	"net"
	"shared"
	"time"
)

type ServerInvoker struct{}

func (ServerInvoker) Invoke() {
	//server request handler

	//servidor de nomes
	marshallerImpl := marshaller.Marshaller{}
	myProxy := auxiliar.Proxy{"locahost", 3300, 1, "server"}
	//registro no servidor de nomes
	reqMrshl, err := json.Marshal(auxiliar.RequestNaming{"Register", "Hasher", myProxy})

	shared.ChecaErro(err, "invoker.go")

	//dial no servidor de nomes
	conn, err := net.DialTimeout("tcp", "localhost:3315", 3*time.Second)

	shared.ChecaErro(err, "invoker.go")

	_, err = conn.Write(reqMrshl)

	shared.ChecaErro(err, "invoker.go")

	conn.Close()

	srhImpl := srh.SRH{ServerHost: "localhost", ServerPort: shared.ServerPort}
	miopPacketReply := miop.Packet{}
	replParams := make([]interface{}, 1)

	hashingPool := pooling.Pool{}
	hashingPool = hashingPool.AllocatePool(shared.POOL_SIZE)

	//hashingImpl := hashing.Hash{}
	calculatorImpl := calculator.Calc{}

	for {
		fmt.Println("aqui")
		rcvMsgBytes := srhImpl.Receive()
		// fmt.Println("invoker", rcvMsgBytes)
		miopPacketRequest := marshallerImpl.Unmarshall(rcvMsgBytes)
		operation := miopPacketRequest.Bd.ReqHeader.Operation

		switch operation {
		case "Hash":
			//acessar a próxima pool que available = true
			availableIndex := 0
			for i := 0; i < shared.POOL_SIZE; i++ {
				if hashingPool.Instances[i].Available == true {
					availableIndex = i
					hashingPool.Instances[i].Available = false
					break
				}
			}
			//fazer essa pool ficar = false
			//fazer a invocação nesse objeto remoto
			//quando acabar, fazer esse objeto remoto ficar available = true
			//fazer isso em alguma goroutine paralelizado.
			fmt.Println("Received message:", miopPacketRequest.Bd.ReqBody.Body[0].(string))
			replParams[0] = hashingPool.Instances[availableIndex].HashSha256(miopPacketRequest.Bd.ReqBody.Body[0].(string))
			fmt.Println("Hashed message:", replParams[0])
			hashingPool.Instances[availableIndex].Available = true
		case "Add":
			p1 := int(miopPacketRequest.Bd.ReqBody.Body[0].(float64))
			p2 := int(miopPacketRequest.Bd.ReqBody.Body[1].(float64))
			fmt.Println("Received nums: ", p1, p2)
			replParams[0] = calculatorImpl.Add(p1, p2)
			fmt.Println("Result: ", replParams[0])
		}

		repHeader := miop.ReplyHeader{Context: "", RequestID: miopPacketRequest.Bd.ReqHeader.RequestID, Status: 1}
		repBody := miop.ReplyBody{OperationResult: replParams}
		header := miop.Header{Magic: "MIOP", Version: "1.0", ByteOrder: true, MessageType: 2}
		body := miop.Body{RepHeader: repHeader, RepBody: repBody}
		miopPacketReply = miop.Packet{Hdr: header, Bd: body}

		msgToClientBytes := marshallerImpl.Marshall(miopPacketReply)
		srhImpl.Send(msgToClientBytes)
	}
}
