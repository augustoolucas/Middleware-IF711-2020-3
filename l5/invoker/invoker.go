package invoker

import (
	"Middleware-IF711-2020-3/l5/calculator"
	"Middleware-IF711-2020-3/l5/hashing"
	"Middleware-IF711-2020-3/l5/marshaller"
	"Middleware-IF711-2020-3/l5/miop"
	"Middleware-IF711-2020-3/l5/srh"
	"shared"
	"fmt"
)

type ServerInvoker struct{}

func (ServerInvoker) Invoke() {
	//server request handler
	srhImpl := srh.SRH{ServerHost: "localhost", ServerPort: shared.CALCULATOR_PORT}
	marshallerImpl := marshaller.Marshaller{}
	miopPacketReply := miop.Packet{}
	replParams := make([]interface{}, 1)
	hashingImpl := hashing.Hash{}
	calculatorImpl := calculator.Calc{}

	for {
		rcvMsgBytes := srhImpl.Receive()

		miopPacketRequest := marshallerImpl.Unmarshall(rcvMsgBytes)
		operation := miopPacketRequest.Bd.ReqHeader.Operation

		switch operation {
		case "Hash":
			fmt.Println("Received message:", miopPacketRequest.Bd.ReqBody.Body[0].(string))
			replParams[0] = hashingImpl.Hashing(miopPacketRequest.Bd.ReqBody.Body[0].(string))
			fmt.Println("Hashed message:", replParams[0])
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
