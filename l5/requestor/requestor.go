package requestor

import (
	"Middleware-IF711-2020-3/l5/auxiliar"
	"Middleware-IF711-2020-3/l5/crh"
	"Middleware-IF711-2020-3/l5/marshaller"
	"shared"

	"Middleware-IF711-2020-3/l5/miop"
)

type Requestor struct{}

func (Requestor) Invoke(inv auxiliar.Invocation) interface{} {
	transportProtocol := shared.TRANSPORT_PROTOCOL
	marshallerInst := marshaller.Marshaller{}
	crhInst := crh.CRH{ServerHost: inv.Host, ServerPort: inv.Port}

	// create request packet
	reqHeader := miop.RequestHeader{Context: "Context", RequestID: 1000, ResponseExpected: true, ObjectKey: 2000, Operation: inv.Request.Op}
	reqBody := miop.RequestBody{Body: inv.Request.Params}
	header := miop.Header{Magic: "MIOP", Version: "1.0", ByteOrder: true, MessageType: 1}
	body := miop.Body{ReqHeader: reqHeader, ReqBody: reqBody}
	miopPacketRequest := miop.Packet{Hdr: header, Bd: body}

	// serialise request packet
	msgToClientBytes := marshallerInst.Marshall(miopPacketRequest)

	// send request packet and receive reply packet
	msgFromServerBytes := crhInst.SendReceive(msgToClientBytes, transportProtocol)
	miopPacketReply := marshallerInst.Unmarshall(msgFromServerBytes)

	// extract result from reply packet
	r := miopPacketReply.Bd.RepBody.OperationResult

	return r
}
