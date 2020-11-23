package hashing

import (
	"../crh"
	"../marshaller"
	"../miop"
	"errors"
	"shared"
)

//Response em sha256
type Response struct {
	PwSha256 string
}

//Request pro hasher
type Request struct {
	Op     string
	Params []interface{}
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
	params := make([]interface{}, 1)
	params[0] = message
	request := Request{Op: "hash", Params: params}
	inv := Invocation{Host: proxy.Host, Port: proxy.Port, Request: request}

	// invoke requestor
	// Invoca o Requestor e aguarda resposta
	req := Requestor{}
	response := req.Invoke(inv, transportProtocol).([]interface{})

	// Envia resposta ao Cliente
	return string(response[0].(string)), nil
}

func (Requestor) Invoke(inv Invocation, transportProtocol string) interface{} {
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
